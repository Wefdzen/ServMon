package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Wefdzen/ServMon/pkg/db/database"
	"github.com/Wefdzen/ServMon/pkg/db/model"
	"github.com/Wefdzen/ServMon/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func DrawAllParam(c *gin.Context) {
	// Get the numServ parameter from the URL (e.g., "/test/1")
	numServ := c.Param("numServ")

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
	tmp := database.GetRecordByIp(userRepo, data[serverIndex].IpServer)

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
		loadAvgCore(tmp),
		ramAvg(tmp),
		memoryAvg(tmp),
	)
	// Render the line chart into the file
	page.Render(io.MultiWriter(f))

	// Return the rendered HTML to the client
	c.HTML(http.StatusOK, "line.html", gin.H{
		"title": "LoadAvg",
	})

}
func loadAvgCore(tmp []model.RecordAboutServerInfo) *charts.Line {
	// Create a new line chart object
	line := charts.NewLine()
	// Set global options for the chart (e.g., Title and Subtitle)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Load avg"}),
	)
	// array time.Now() - 5min ~ time.Now()
	var xLabels []string
	currentTime := time.Now()
	for i := 12; i >= 0; i-- {
		timeLabel := currentTime.Add(-time.Duration(i) * time.Minute * 5).Format("15:04")
		xLabels = append(xLabels, timeLabel)
	}
	xLabels = reverseSliceString(xLabels)

	// Set the X-axis to represent 24 hours and add the series data for this server
	line.SetXAxis(xLabels).
		AddSeries("Load avg", GenerateLineItems(tmp)).
		//See value
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: opts.Bool(true),
			}),
			charts.WithLabelOpts(opts.Label{
				Show: opts.Bool(true),
			}),
		)
	//settgings
	line.SetGlobalOptions(
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Load averages", // Name of Axis Y
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Time",
		}),
	)
	return line
}

func ramAvg(tmp []model.RecordAboutServerInfo) *charts.Line {
	// Create a new line chart object
	line := charts.NewLine()
	// Set global options for the chart (e.g., Title and Subtitle)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "RAM avg"}),
	)

	var xLabels []string
	currentTime := time.Now()
	for i := 12; i >= 0; i-- {
		timeLabel := currentTime.Add(-time.Duration(i) * time.Minute * 5).Format("15:04")
		xLabels = append(xLabels, timeLabel)
	}
	xLabels = reverseSliceString(xLabels)

	// Set the X-axis to represent 24 hours and add the series data for this server
	line.SetXAxis(xLabels).
		AddSeries("Ram avg", GenerateLineRam(tmp)).
		//See value
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: opts.Bool(true),
			}),
			charts.WithLabelOpts(opts.Label{
				Show: opts.Bool(true),
			}),
		)
	//settgings
	line.SetGlobalOptions(
		charts.WithYAxisOpts(opts.YAxis{
			Name: "RAM (MB)",     // Name of Axis Y
			Max:  GetMaxRam(tmp), // set max value for axis Y
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Time",
		}),
	)
	return line
}

func memoryAvg(tmp []model.RecordAboutServerInfo) *charts.Line {
	// Create a new line chart object
	line := charts.NewLine()
	// Set global options for the chart (e.g., Title and Subtitle)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Memory avg"}),
	)

	var xLabels []string
	currentTime := time.Now()
	for i := 12; i >= 0; i-- {
		timeLabel := currentTime.Add(-time.Duration(i) * time.Minute * 5).Format("15:04")
		xLabels = append(xLabels, timeLabel)
	}
	xLabels = reverseSliceString(xLabels)

	// Set the X-axis to represent 24 hours and add the series data for this server
	line.SetXAxis(xLabels).
		AddSeries("Memory avg", GenerateLineMemory(tmp)).
		//See value
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: opts.Bool(true),
			}),
			charts.WithLabelOpts(opts.Label{
				Show: opts.Bool(true),
			}),
		)
	//settgings
	line.SetGlobalOptions(
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Memory (GB)",     // Name of Axis Y
			Max:  GetMaxMemory(tmp), //set max value for axis Y
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Time",
		}),
	)
	return line
}

func reverseSliceString(input []string) []string {
	output := make([]string, len(input))
	j := 0
	for i := len(input) - 1; i >= 0; i-- {
		output[j] = input[i]
		j++
	}
	return output
}
