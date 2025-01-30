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
func RecordDataServerToFile(servers []models.Server) error {
	data, err := json.MarshalIndent(servers, "", "	")
	if err != nil {
		fmt.Println("marshal error")
		return err
	}

	file, err := os.Create("./servers.json")
	if err != nil {
		log.Fatalf("open file error %v\n", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("write file error %v\n", err)
		return err
	}

	fmt.Println("Data successfully written to servers.json")
	return nil
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
		return 0, err
	}
	return config.CountServersOfUser, nil
}

func GetInfoServers(fileName string) ([]models.Server, error) {
	data, err := os.ReadFile(fileName) // Read file
	if err != nil {
		log.Fatalf("Error read from file: %v\n", err)
		return nil, err
	}
	// Get data from file
	var dataAboutServers []models.Server
	err = json.Unmarshal(data, &dataAboutServers)
	if err != nil {
		log.Fatalf("Error unmarshal file: %v\n", err)
		return nil, err
	}
	return dataAboutServers, nil
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
