package launcapp

import (
	"fmt"

	inituser "github.com/Wefdzen/ServMon/pkg/initUser"
	"github.com/Wefdzen/ServMon/pkg/service"
)

// LaunchApp responsible for launc main app.
func LaunchApp() {
	// Get data of servers
	servers, err := inituser.GetServersData()
	if err != nil {
		fmt.Println("Error can't init Servers")
		return
	}

	// Record data about serevers to file
	service.RecordDataServerToFile(servers)
}
