package launcapp

import (
	"encoding/json"
	"fmt"
	"os"

	inituser "github.com/Wefdzen/ServMon/pkg/initUser"
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
	data, err := json.MarshalIndent(servers, "", "	")
	if err != nil {
		fmt.Println("marshal error")
		return
	}
	file, err := os.OpenFile("servers.json", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file error")
		return
	}
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("write file error")
		return
	}

	fmt.Println("Data successfully written to servers.json")
}
