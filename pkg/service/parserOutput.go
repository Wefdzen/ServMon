package service

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"

	"github.com/Wefdzen/ServMon/pkg/models"
)

func ParseSystemStats(output string, res models.Server) models.ServerInfo {
	var info models.ServerInfo
	info.NameOfService = res.NameOfService
	info.IpServer = res.IpServer

	info.CoreCount = uint8(runtime.NumCPU())
	//get for 5min (load average: X.X)
	loadAvgRegex := regexp.MustCompile(`load average:\s*([\d\.]+),\s*([\d\.]+),\s*([\d\.]+)`)
	loadAvgMatches := loadAvgRegex.FindStringSubmatch(output)
	if len(loadAvgMatches) > 1 {
		info.LoadAvg5Min = loadAvgMatches[1]
	}

	// get RAM (Mem: total used free shared buff/cache available)
	memRegex := regexp.MustCompile(`Mem:\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)
	memMatches := memRegex.FindStringSubmatch(output)
	if len(memMatches) > 6 {
		totalMem, _ := strconv.Atoi(memMatches[1])
		usedMem, _ := strconv.Atoi(memMatches[2])
		info.Ram = fmt.Sprintf("%d/%d MB", usedMem, totalMem)
	}
	// get memory
	diskRegex := regexp.MustCompile(`/dev/\S+\s+([\d\.]+)[GM]+\s+([\d\.]+)[GM]+\s+([\d\.]+)[GM]+`)
	diskMatches := diskRegex.FindStringSubmatch(output)
	if len(diskMatches) > 3 {
		usedSize, _ := strconv.ParseFloat(diskMatches[2], 64)
		totalSize, _ := strconv.ParseFloat(diskMatches[1], 64)
		info.Memory = fmt.Sprintf("Used %.1f GB of %.1f GB", usedSize, totalSize)
	}

	return info
}
