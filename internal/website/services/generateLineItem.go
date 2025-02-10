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
