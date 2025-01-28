package inituser

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Wefdzen/ServMon/pkg/models"
	"github.com/Wefdzen/ServMon/pkg/service"
	"github.com/charmbracelet/huh"
)

// AddNewServer add new server to file
func AddNewServer(newServ models.Server, fileName string) error {
	data, err := os.ReadFile(fileName) // Read file
	if err != nil {
		log.Fatalf("Error read from file: %v\n", err)
		return err
	}
	// Get data from file
	var servers []models.Server
	err = json.Unmarshal(data, &servers)
	if err != nil {
		log.Fatalf("Error unmarshal file: %v\n", err)
		return err
	}

	//Add new recored
	servers = append(servers, newServ)

	//Overwrite file with new server
	data, err = json.MarshalIndent(servers, "", "	")
	if err != nil {
		log.Fatalf("marshal error %v\n", err)
		return err
	}
	file, err := os.Create(fileName)
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
	//Get counter
	tmp, _ := service.GetCountServer("./pkg/config/config.json")
	tmp++
	//Set new counter
	service.SetNewCountServerConfig("./pkg/config/config.json", tmp)
	return nil
}

// DeleteNewServer delete info about server from file TODO(loss id)
func DeleteNewServer(ip, fileName string) error {
	data, err := os.ReadFile(fileName) // Read file
	if err != nil {
		log.Fatalf("Error read from file: %v\n", err)
		return err
	}
	// Get data from file
	var servers []models.Server
	err = json.Unmarshal(data, &servers)
	if err != nil {
		log.Fatalf("Error unmarshal file: %v\n", err)
		return err
	}

	//delete server by id
	for i := range servers {
		if servers[i].IpServer == ip {
			servers = append(servers[:i], servers[i+1:]...)
		}
	}

	//Overwrite file without server
	data, err = json.MarshalIndent(servers, "", "	")
	if err != nil {
		log.Fatalf("marshal error %v\n", err)
		return err
	}
	file, err := os.Create(fileName)
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
	//Get counter
	tmp, _ := service.GetCountServer("./pkg/config/config.json")
	tmp--
	//Set new counter
	service.SetNewCountServerConfig("./pkg/config/config.json", tmp)
	return nil
}

// GetDataAboutNewServer form huh for 1 server
func GetDataAboutNewServer() models.Server {
	serverData := models.Server{}
	huh.NewInput().
		Title("Login(exm: root)").
		Prompt("?").
		Value(&serverData.Account).
		Run()
	huh.NewInput().
		Title("Ip of server").
		Prompt("?").
		Value(&serverData.IpServer).
		Run()
	huh.NewInput().
		Title("Password of server").
		Prompt("?").
		Value(&serverData.Password).
		Run()
	//Success
	return serverData
}
