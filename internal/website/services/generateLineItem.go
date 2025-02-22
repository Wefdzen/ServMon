package services

import (
	"github.com/Wefdzen/ServMon/pkg/db/model"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GenerateLineItems(res []model.RecordAboutServerInfo) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(res); i++ {
		items = append(items, opts.LineData{Value: res[i].LoadAvg5Min})
	}
	return items
}

func GenerateLineRam(res []model.RecordAboutServerInfo) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(res); i++ {
		curRam, _, _ := parseRam(res[i].Ram)
		items = append(items, opts.LineData{Value: curRam})
	}
	return items
}

func GenerateLineMemory(res []model.RecordAboutServerInfo) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(res); i++ {
		curMemory, _, _ := parseMemory(res[i].Memory)
		items = append(items, opts.LineData{Value: curMemory})
	}
	return items
}

func GetMaxMemory(res []model.RecordAboutServerInfo) string {
	if len(res) <= 0 {
		return "haven't"
	}
	_, maxMemory, _ := parseMemory(res[0].Memory)
	return maxMemory
}

func GetMaxRam(res []model.RecordAboutServerInfo) string {
	if len(res) <= 0 {
		return "haven't"
	}
	_, maxRam, _ := parseRam(res[0].Ram)
	return maxRam
}
