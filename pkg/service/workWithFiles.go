package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Wefdzen/ServMon/pkg/config"
	"github.com/Wefdzen/ServMon/pkg/models"
)

// RecordDataServerToFile record data about servers to file only for init user add here don't work.
func RecordDataServerToFile(servers []models.Server) {
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
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("write file error")
		return
	}

	fmt.Println("Data successfully written to servers.json")

}

// GetCountServer get count of server now
func GetCountServer(fileName string) (uint8, error) {
	data, err := os.ReadFile(fileName) // Read file
	if err != nil {
		log.Fatalf("Error read from file: %v\n", err)
		return 0, err
	}
	// Get data from file
	var config config.ConfigUser
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshal file: %v\n", err)
		return 0, nil
	}
	return config.CountServersOfUser, nil
}

// SetNewCountServerConfig set new count servers in config file.
func SetNewCountServerConfig(fileName string, newCountServers uint8) {
	data, err := os.ReadFile(fileName) // Read file
	if err != nil {
		log.Fatalf("Error read from file: %v\n", err)
		return
	}
	// Get data from file
	var config config.ConfigUser
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshal file: %v\n", err)
		return
	}

	//upate count servers
	config.CountServersOfUser = newCountServers

	//Overwrite file with new count servers
	data, err = json.MarshalIndent(config, "", "	")
	if err != nil {
		fmt.Println("marshal error")
		return
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("write file error")
		return
	}

}
