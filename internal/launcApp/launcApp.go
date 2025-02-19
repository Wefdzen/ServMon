package launcapp

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
			// sigChan := make(chan os.Signal, 1)
			// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

			// done := make(chan bool, 1)

			// go func() {
			// 	sig := <-sigChan // wait a signal
			// 	fmt.Printf("\nSignal received: %v\n", sig)
			// 	// Clear file for future use those data will be not relevant
			// 	file, err := os.Create("./internal/launchApp/lastRecord.json")
			// 	if err != nil {
			// 		fmt.Println(err)
			// 	}
			// 	defer file.Close()

			// 	done <- true // Send signal end
			// }()
			go case5()
		case "6":
			fmt.Println("here will be on/off of smtp")
			case6()
		}
	}
}

func case5() {
	var records []model.RecordAboutServerInfo
	for {
		data, err := service.GetInfoServers("./servers.json")
		if err != nil {
			fmt.Println(err)
			return
		}

		// get command from bash file
		myCmd, err := os.ReadFile("./pkg/workWithServers/commandToServer.bash")
		if err != nil {
			fmt.Print(err)
		}

		// launch command in servers, len(data) is count of servers
		for i := 0; i < len(data); i++ {
			output, err := workwithservers.SendCommandToServer(data[i].IpServer, data[i].Account, data[i].Password, string(myCmd), "")
			if err != nil {
				fmt.Println(err)
				return
			}
			mdl := service.ParseSystemStats(output, data[i])

			// create record
			rec := model.RecordAboutServerInfo{
				Time:        time.Now().Unix(),
				NameService: mdl.NameOfService,
				IpServer:    mdl.IpServer,
				LoadAvg5Min: mdl.LoadAvg5Min,
				Ram:         mdl.Ram,
				Memory:      mdl.Memory,
			}

			// add record to slice
			records = append(records, rec)

			//work with file
			file, err := os.Create("./internal/launcApp/lastRecord.json")
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			dataJson, err := json.MarshalIndent(records, "", "	")
			if err != nil {
				fmt.Println(err)
			}
			_, err = file.Write(dataJson)
			if err != nil {
				fmt.Println(err)
			}

			// if we have 12 records, calculate average and save to db
			if len(records) == 12 {
				avgRecord := calculateAverage(records)
				userRepo := database.NewGormUserRepository()
				database.AddNewRecord(userRepo, avgRecord)
				records = records[:0] // reset records slice
			}
		}
		time.Sleep(1 * time.Minute) // every 5 min will be
	}
}

func calculateAverage(records []model.RecordAboutServerInfo) *model.RecordAboutServerInfo {
	var sumLoadAvg, sumRam float64
	for _, rec := range records {
		loadAvg, _ := strconv.ParseFloat(rec.LoadAvg5Min, 64)
		ram, _ := strconv.ParseFloat(rec.Ram, 64)
		sumLoadAvg += loadAvg
		sumRam += ram
	}

	avgRecord := records[0]
	avgRecord.LoadAvg5Min = fmt.Sprintf("%v", sumLoadAvg/float64(len(records)))
	avgRecord.Ram = fmt.Sprintf("%v", sumRam/float64(len(records)))
	avgRecord.Memory = records[len(records)-1].Memory // use last 'cause last record is True

	return &avgRecord
}

func case6() {
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

}

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
			huh.NewOption("6. On/Off smtp mailout about crit loadAvg", "6"),
		).
		Value(&mode).
		Run()
	return mode
}
