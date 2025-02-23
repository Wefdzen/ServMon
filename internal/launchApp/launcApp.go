package launchcapp

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/huh"

	"github.com/Wefdzen/ServMon/internal/website/services"
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
			fmt.Println("here will be on/off of smtp NOW NOT WORK")
			case6()
		}
	}
}
func init() {
	//clear lastRecord.json file
	file, err := os.Create("./internal/launchApp/lastRecord.json")
	var clear []int
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		dataJson, err := json.MarshalIndent(clear, "", "  ") // Записываем массив, а не map
		if err != nil {
			fmt.Println(err)
		}
		_, err = file.Write(dataJson)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func case5() {
	var recordsByServer = make(map[string][]model.RecordAboutServerInfo) // Ключ — IP сервера

	for {
		data, err := service.GetInfoServers("./servers.json")
		if err != nil {
			fmt.Println(err)
			return
		}

		myCmd, err := os.ReadFile("./pkg/workWithServers/commandToServer.bash")
		if err != nil {
			fmt.Print(err)
		}

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

			// Добавляем в мапу записей для конкретного сервера
			recordsByServer[rec.IpServer] = append(recordsByServer[rec.IpServer], rec)

			if len(recordsByServer[rec.IpServer]) >= 24 {
				recordsByServer[rec.IpServer] = recordsByServer[rec.IpServer][12:]
			}
			// Проверяем, есть ли 12 записей для текущего сервера
			if len(recordsByServer[rec.IpServer]) == 12 { //13 уже не надо ждем 24 он очищает => снова получается 12 и оно тут проходит
				// if we get 12 rec => if first of 12 rec is old by 1 hour 4 min(just in case) this data is old and we need delete this
				lastRecord := recordsByServer[rec.IpServer]
				if lastRecord[0].Time > int64(time.Now().Add(-(time.Hour + (time.Minute * 4))).Unix()) {
					avgRecord := calculateAverage(recordsByServer[rec.IpServer]) // Считаем среднее
					userRepo := database.NewGormUserRepository()
					database.AddNewRecord(userRepo, avgRecord)
				} else { // clear file 'cause we need new data
					recordsByServer[rec.IpServer] = recordsByServer[rec.IpServer][0:0]
				}

			}
		}

		// Подготовка массива всех записей перед записью в файл
		var allRecords []model.RecordAboutServerInfo
		for _, records := range recordsByServer {
			allRecords = append(allRecords, records...) // Добавляем все записи из каждого сервера
		}

		// Записываем в JSON только массив
		file, err := os.Create("./internal/launchApp/lastRecord.json")
		if err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()
			dataJson, err := json.MarshalIndent(allRecords, "", "  ") // Записываем массив, а не map
			if err != nil {
				fmt.Println(err)
			}
			_, err = file.Write(dataJson)
			if err != nil {
				fmt.Println(err)
			}
		}

		time.Sleep(5 * time.Minute)
	}
}

func calculateAverage(records []model.RecordAboutServerInfo) *model.RecordAboutServerInfo {
	var sumLoadAvg float64
	var sumRam int
	for _, rec := range records {
		loadAvg, _ := strconv.ParseFloat(rec.LoadAvg5Min, 64)
		ram, _, _ := services.ParseRam(rec.Ram)
		tmp, _ := strconv.Atoi(ram)
		sumLoadAvg += loadAvg
		sumRam += tmp
	}

	avgRecord := records[0]
	avgRecord.LoadAvg5Min = fmt.Sprintf("%.2f", sumLoadAvg/float64(len(records)))

	_, max, _ := services.ParseRam(avgRecord.Ram)
	avgRecord.Ram = fmt.Sprintf("%v/%v MB", sumRam/len(records), max)
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
			huh.NewOption("1. Launch the app for the first time", "1"),
			huh.NewOption("2. Add a new server", "2"),
			huh.NewOption("3. Delete a server", "3"),
			huh.NewOption("4. Exit", "4"),
			huh.NewOption("5. Send command (start record data)", "5"),
			huh.NewOption("6. Toggle SMTP notifications for critical load average", "6"),
		).
		Value(&mode).
		Run()
	return mode
}
