package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Wefdzen/ServMon/pkg/db/database"
	"github.com/Wefdzen/ServMon/pkg/db/model"
	"github.com/Wefdzen/ServMon/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/components"
)

func DrawAllParamWithMode(c *gin.Context) {
	// Get the numServ parameter from the URL (e.g., "/test/1")
	numServ := c.Param("numServ")
	modeFromParam := c.Param("mode")
	mode, _ := strconv.Atoi(modeFromParam)
	// Get the user repository
	userRepo := database.NewGormUserRepository()

	// Read server data from the JSON file
	data, err := service.GetInfoServers("./servers.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if a server with the specified numServ exists
	// Assume numServ is an index or a server number
	serverIndex := -1
	for i := range data {
		if fmt.Sprintf("%d", i+1) == numServ {
			serverIndex = i
			break
		}
	}

	// If no server is found with the given numServ, return an error
	if serverIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
		return
	}

	// Get the data for the specific server
	var tmp []model.RecordAboutServerInfo
	if mode != 1 {
		tmp = database.GetRecordByIp(userRepo, data[serverIndex].IpServer)
	} else { // if man choose 1 hour mode => need use not db -> tmp file with 12 record every 5min but he update every 1hour
		dataOfLastRecord, err := os.ReadFile("./internal/launcApp/lastRecord.json") // Read file
		if err != nil {
			fmt.Println(err)
		}
		var SortServ []model.RecordAboutServerInfo
		err = json.Unmarshal(dataOfLastRecord, &SortServ)
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < len(SortServ); i++ {
			if SortServ[i].IpServer == data[serverIndex].IpServer {
				tmp = append(tmp, SortServ[i])
			}
		}
	}

	// Create a file for rendering the chart
	f, err := os.Create("./internal/website/static/html/line.html")
	if err != nil {
		// Handle error if the file cannot be created
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}
	defer f.Close()
	page := components.NewPage()
	page.AddCharts(
		loadAvgCore(tmp, mode),
		ramAvg(tmp, mode),
		memoryAvg(tmp, mode),
	)
	// Render the line chart into the file
	page.Render(io.MultiWriter(f))

	// Return the rendered HTML to the client
	c.HTML(http.StatusOK, "test.html", gin.H{
		"title": "LoadAvg",
	})

}
