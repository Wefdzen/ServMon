package services

import (
	"fmt"
	"strconv"

	"github.com/Wefdzen/ServMon/pkg/db/model"
)

func PlusAvg(rec1 model.RecordAboutServerInfo, rec2 model.RecordAboutServerInfo) model.RecordAboutServerInfo {
	var tmp model.RecordAboutServerInfo
	tmp.IpServer = rec1.IpServer
	tmp.Time = rec2.Time //ok
	num1, _ := strconv.ParseFloat(rec1.LoadAvg5Min, 64)
	num2, _ := strconv.ParseFloat(rec2.LoadAvg5Min, 64)
	tmp.LoadAvg5Min = fmt.Sprintf("%.2f", (num1+num2)/2)

	cur1, max1, _ := ParseRam(rec1.Ram)
	cur2, _, _ := ParseRam(rec2.Ram)
	cur1Ram, _ := strconv.Atoi(cur1)
	cur2Ram, _ := strconv.Atoi(cur2)
	tmp.Ram = fmt.Sprintf("%v/%v MB", (cur1Ram+cur2Ram)/2, max1)

	tmp.NameService = rec1.NameService
	tmp.Memory = rec1.Memory
	return tmp
}
