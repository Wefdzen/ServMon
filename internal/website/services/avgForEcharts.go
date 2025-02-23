package services

import (
	"time"

	"github.com/Wefdzen/ServMon/pkg/db/model"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// loadAvgCore return line for echarts. Create averages of Core, mode=> 1 is 1 last hour info, 2 is 12 hours info,
// 3 is 24 hours(1 day)
func loadAvgCore(tmp []model.RecordAboutServerInfo, mode int) *charts.Line {
	// Create a new line chart object
	line := charts.NewLine()
	// Set global options for the chart (e.g., Title and Subtitle)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Load avg"}),
	)

	var xLabels []string
	// array time.Now() - 5min ~ time.Now()
	switch mode {
	case 1:
		currentTime := time.Now()
		for i := 12; i >= 0; i-- {
			timeLabel := currentTime.Add(-time.Duration(i) * time.Minute * 5).Format("15:04")
			xLabels = append(xLabels, timeLabel)
		}
	case 2:
		currentTime := time.Now()
		for i := 12; i >= 0; i-- {
			timeLabel := currentTime.Add(-time.Duration(i) * time.Hour).Format("15:04")
			xLabels = append(xLabels, timeLabel)
		}
	case 3:
		currentTime := time.Now()
		for i := 12; i >= 0; i-- {
			timeLabel := currentTime.Add(-time.Duration(i) * time.Hour * 2).Format("15:04")
			xLabels = append(xLabels, timeLabel)
		}
	}
	xLabels = reverseSliceString(xLabels)

	// Set the X-axis to represent 24 hours and add the series data for this server
	if mode == 1 {
		tmp = reverseRecordAboutServer(tmp)
	}
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
	//settings
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

// ramAvg return line for echarts. Create averages of Ram, mode=> 1 is 1 last hour info, 2 is 12 hours info,
// 3 is 24 hours(1 day)
func ramAvg(tmp []model.RecordAboutServerInfo, mode int) *charts.Line {
	// Create a new line chart object
	line := charts.NewLine()
	// Set global options for the chart (e.g., Title and Subtitle)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "RAM avg"}),
	)

	var xLabels []string
	switch mode {
	case 1:
		currentTime := time.Now()
		for i := 12; i >= 0; i-- {
			timeLabel := currentTime.Add(-time.Duration(i) * time.Minute * 5).Format("15:04")
			xLabels = append(xLabels, timeLabel)
		}
	case 2:
		currentTime := time.Now()
		for i := 12; i >= 0; i-- {
			timeLabel := currentTime.Add(-time.Duration(i) * time.Hour).Format("15:04")
			xLabels = append(xLabels, timeLabel)
		}
	case 3:
		currentTime := time.Now()
		for i := 12; i >= 0; i-- {
			timeLabel := currentTime.Add(-time.Duration(i) * time.Hour * 2).Format("15:04")
			xLabels = append(xLabels, timeLabel)
		}
	}
	xLabels = reverseSliceString(xLabels)

	// Set the X-axis to represent 24 hours and add the series data for this server
	if mode == 1 {
		tmp = reverseRecordAboutServer(tmp)
	}
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

// memoryAvg return line for echarts. Create averages of Core, mode=> 1 is 1 last hour info, 2 is 12 hours info,
// 3 is 24 hours(1 day)
func memoryAvg(tmp []model.RecordAboutServerInfo, mode int) *charts.Line {
	// Create a new line chart object
	line := charts.NewLine()
	// Set global options for the chart (e.g., Title and Subtitle)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Memory avg"}),
	)

	var xLabels []string
	switch mode {
	case 1:
		currentTime := time.Now()
		for i := 12; i >= 0; i-- {
			timeLabel := currentTime.Add(-time.Duration(i) * time.Minute * 5).Format("15:04")
			xLabels = append(xLabels, timeLabel)
		}
	case 2:
		currentTime := time.Now()
		for i := 12; i >= 0; i-- {
			timeLabel := currentTime.Add(-time.Duration(i) * time.Hour).Format("15:04")
			xLabels = append(xLabels, timeLabel)
		}
	case 3:
		currentTime := time.Now()
		for i := 12; i >= 0; i-- {
			timeLabel := currentTime.Add(-time.Duration(i) * time.Hour * 2).Format("15:04")
			xLabels = append(xLabels, timeLabel)
		}
	}
	xLabels = reverseSliceString(xLabels)

	// Set the X-axis to represent 24 hours and add the series data for this server
	if mode == 1 {
		tmp = reverseRecordAboutServer(tmp)
	}
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

func reverseRecordAboutServer(input []model.RecordAboutServerInfo) []model.RecordAboutServerInfo {
	output := make([]model.RecordAboutServerInfo, len(input))
	j := 0
	for i := len(input) - 1; i >= 0; i-- {
		output[j] = input[i]
		j++
	}
	return output
}
