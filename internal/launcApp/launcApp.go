package launcapp

import (
	"fmt"

	"github.com/charmbracelet/huh"

	inituser "github.com/Wefdzen/ServMon/pkg/initUser"
	"github.com/Wefdzen/ServMon/pkg/service"
	workwithservers "github.com/Wefdzen/ServMon/pkg/workWithServers"
)

// LaunchApp responsible for launc main app.
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
			service.RecordDataServerToFile(servers)
		case "2":
			newServer := inituser.GetDataAboutNewServer()
			inituser.AddNewServer(newServer, "./servers.json")
		case "3":
			inituser.DeleteNewServer(inituser.GetIpOfServer(), "./servers.json")
		case "4":
			fmt.Println("end settings")
			return
		case "5":
			data, err := service.GetInfoServers("./servers.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			for i := 0; i < len(data); i++ {
				output, err := workwithservers.SendCommandToServer(data[i].IpServer, data[i].Account, data[i].Password, "free", "")
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(output)
			}
		}

	}
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
			huh.NewOption("5. Send command free", "5"),
		).
		Value(&mode). //
		Run()
	return mode
}
