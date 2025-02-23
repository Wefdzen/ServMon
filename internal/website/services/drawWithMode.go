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
	// Get the numServ parameter from the URL (e.g., "/test/1/1")
	numServ := c.Param("numServ")
	modeFromParam := c.Param("mode")
	mode, _ := strconv.Atoi(modeFromParam)
	//validate mode
	if mode < 1 || mode > 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Uncorrect mode",
		})
	}
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
	switch mode {
	case 1: // если пользователь выбрал режим 1 час, нужно использовать файл с 12 записями каждые 5 минут
		dataOfLastRecord, err := os.ReadFile("./internal/launchApp/lastRecord.json") // Чтение файла
		if err != nil {
			fmt.Println(err)
		}

		var SortServ []model.RecordAboutServerInfo
		err = json.Unmarshal(dataOfLastRecord, &SortServ)
		if err != nil {
			fmt.Println(err)
		}

		// Фильтруем записи по IP сервера
		for _, record := range SortServ {
			if record.IpServer == data[serverIndex].IpServer {
				tmp = append(tmp, record)
			}
		}

		// Если записей больше 12, берем только последние 12
		if len(tmp) > 12 {
			tmp = tmp[len(tmp)-12:] // Берем последние 12 записей
		}
	case 2: //12hours
		tmp = database.GetRecordByIp(userRepo, data[serverIndex].IpServer, 12)
	case 3: //24hours
		tmp = database.GetRecordByIp(userRepo, data[serverIndex].IpServer, 24)
		//get avg by 2 rec to 1rec
		if len(tmp) == 0 || len(tmp) == 1 {
			break
		}

		var rec24 []model.RecordAboutServerInfo

		if len(tmp)%2 == 0 { //2, 4, 6, 8, 24rec
			for i := 0; i < len(tmp)-1; i += 2 {
				rec24 = append(rec24, PlusAvg(tmp[i], tmp[i+1]))
			}
		} else { //3, 5 , 7, 23 rec
			for i := 0; i < len(tmp)-1; i += 2 {
				if i == len(tmp)-1 { //last element
					rec24 = append(rec24, tmp[i])
					break
				}
				rec24 = append(rec24, PlusAvg(tmp[i], tmp[i+1]))
			}
		}
		tmp = rec24[0:]

		// case 4: // 48 hours
		// 	tmp = database.GetRecordByIp(userRepo, data[serverIndex].IpServer, 48)
		// 	// get avg by 4 rec to 1 rec
		// 	if len(tmp) == 0 || len(tmp) < 4 {
		// 		break
		// 	}

		// 	var rec48 []model.RecordAboutServerInfo

		// 	// Если количество записей кратно 4
		// 	if len(tmp)%4 == 0 {
		// 		for i := 0; i < len(tmp)-3; i += 4 {
		// 			rec48 = append(rec48, PlusAvg(PlusAvg(tmp[i], tmp[i+1]), PlusAvg(tmp[i+2], tmp[i+3])))
		// 		}
		// 	} else { // Если записей не кратно 4
		// 		for i := 0; i < len(tmp)-3; i += 4 {
		// 			rec48 = append(rec48, PlusAvg(PlusAvg(tmp[i], tmp[i+1]), PlusAvg(tmp[i+2], tmp[i+3])))
		// 		}
		// 		// Если остались "лишние" записи, просто добавляем последнюю
		// 		if len(tmp)%4 != 0 {
		// 			rec48 = append(rec48, tmp[len(tmp)-1])
		// 		}
		// 	}

		// 	tmp = rec48[0:]
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
