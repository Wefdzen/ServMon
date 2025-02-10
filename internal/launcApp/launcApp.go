package launcapp

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/huh"

	"github.com/Wefdzen/ServMon/pkg/db/database"
	"github.com/Wefdzen/ServMon/pkg/db/model"
	inituser "github.com/Wefdzen/ServMon/pkg/initUser"
	"github.com/Wefdzen/ServMon/pkg/service"
	workwithservers "github.com/Wefdzen/ServMon/pkg/workWithServers"
)

// LaunchApp responsible for launch main app.
func LaunchApp() {
	for {
		switch menu() {
		case "1":
			// Get data of servers
			servers, err := inituser.GetServersData()
			if err != nil {
				fmt.Println("Error can't init Servers")
				return
			}
			// Record data about servers to file
			err = service.RecordDataServerToFile(servers)
			if err != nil {
				fmt.Println(err)
			}
		case "2":
			newServer := inituser.GetDataAboutNewServer()
			inituser.AddNewServer(newServer, "./servers.json")
		case "3":
			inituser.DeleteNewServer(inituser.GetIpOfServer(), "./servers.json")
		case "4":
			fmt.Println("end settings")
			return
		case "5":
			go case5()
		}

	}
}

func case5() {
	for { // inf
		data, err := service.GetInfoServers("./servers.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		//get command from bash file
		myCmd, err := os.ReadFile("./pkg/workWithServers/commandToServer.bash") // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		//launch command in servers
		for i := 0; i < len(data); i++ {
			output, err := workwithservers.SendCommandToServer(data[i].IpServer, data[i].Account, data[i].Password, string(myCmd), "")
			if err != nil {
				fmt.Println(err)
				return
			}
			mdl := service.ParseSystemStats(output, data[i])
			//TODO mb add smtp if bad
			// statusCode := []string{"Ok", "Can be better", "Bad", "VERY BAD"}
			// tmp, _ := strconv.ParseFloat(mdl.LoadAvg5Min, 64)
			// switch val := tmp / float64(mdl.CoreCount); {
			// case val <= 0.7:
			// 	fmt.Println("Load Avg (5 min):", mdl.LoadAvg5Min, "Status:", statusCode[0])
			// case val <= 2.0:
			// 	fmt.Println("Load Avg (5 min):", mdl.LoadAvg5Min, "Status:", statusCode[1])
			// case val < 5.0:
			// 	fmt.Println("Load Avg (5 min):", mdl.LoadAvg5Min, "Status:", statusCode[2])
			// case val >= 5.0:
			// 	fmt.Println("Load Avg (5 min):", mdl.LoadAvg5Min, "Status:", statusCode[3])
			// default:
			// 	fmt.Println("something wrong :(")
			// }
			// fmt.Println("RAM:", mdl.Ram)
			// fmt.Println("Disk Usage:", mdl.Memory)

			//work with db test
			rec := model.RecordAboutServerInfo{
				Time:        time.Now().Unix(),
				NameService: mdl.NameOfService,
				IpServer:    mdl.IpServer,
				LoadAvg5Min: mdl.LoadAvg5Min,
				Ram:         mdl.Ram,
				Memory:      mdl.Memory,
			}
			userRepo := database.NewGormUserRepository()
			database.AddNewRecord(userRepo, &rec)
		}
		time.Sleep(5 * time.Minute) // every 5 min will be
	}
}

// func getAvg(resultOf []model.RecordAboutServerInfo) *model.RecordAboutServerInfo {
// 	res := resultOf[0]
// 	var sum float64
// 	for i := 0; i < len(resultOf); i++ {
// 		tmp, _ := strconv.ParseFloat(resultOf[i].LoadAvg5Min, 64)
// 		sum += tmp
// 	}
// 	res.LoadAvg5Min = fmt.Sprintf("%v", sum/float64(len(resultOf)))
// 	sum = 0.0
// 	for i := 0; i < len(resultOf); i++ {
// 		tmp, _ := strconv.ParseFloat(resultOf[i].Ram, 64)
// 		sum += tmp
// 	}
// 	res.Ram = fmt.Sprintf("%v", )
// 	res.Memory = resultOf[len(resultOf)-1].Memory //use last 'cause last record is True
// 	return &res
// }

func menu() string {
	var mode string
	huh.NewSelect[string]().
		Title("Select an action").
		Options(
			huh.NewOption("1. If first open this app", "1"),
			huh.NewOption("2. Add new server", "2"),
			huh.NewOption("3. Delete server", "3"),
			huh.NewOption("4. Exit", "4"),
			huh.NewOption("5. Send command(daemon)", "5"),
		).
		Value(&mode).
		Run()
	return mode
}
