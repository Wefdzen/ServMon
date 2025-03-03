package inituser

import (
	"errors"
	"strconv"

	"github.com/Wefdzen/ServMon/pkg/models"
	"github.com/Wefdzen/ServMon/pkg/service"
	"github.com/charmbracelet/huh"
)

// GetServersData launch form for get info about servers.
func GetServersData() ([]models.Server, error) {
	UserServers := make([]models.Server, 0, 3)
	tmp := ""
	//Get count servers
	huh.NewInput().
		Title("How many servers do you want to monitor").
		Prompt("?").
		Value(&tmp).
		Run()

	countServer, err := strconv.Atoi(tmp)
	if err != nil {
		return nil, err
	}
	// Check count of server
	if countServer <= 0 {
		return nil, errors.New("count of server <= 0 :(")
	}

	var idCount uint8 = 0
	//Get data about servers
	for i := 0; i < countServer; i++ {
		serverData := models.Server{}
		huh.NewInput().
			Title("Name of service (e.g.: DigitalOcean)").
			Prompt("?").
			Value(&serverData.NameOfService).
			Run()
		huh.NewInput().
			Title("Login (e.g.: root)").
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

		idCount++
		serverData.Id = idCount
		service.SetNewCountServerConfig("./pkg/config/config.json", idCount)

		//add new server
		UserServers = append(UserServers, serverData)
	}

	return UserServers, nil
}
