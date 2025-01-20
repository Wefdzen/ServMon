package inituser

import (
	"strconv"

	"github.com/charmbracelet/huh"
)

// GetServersData launch form for get info about servers.
func GetServersData() ([]Server, error) {
	UserServers := make([]Server, 0, 3)
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

	//Get data about servers
	for i := 0; i < countServer; i++ {
		serverData := Server{}
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

		//TODOserverData.Id = config.CountServersOfUser
		serverData.Id = uint8(i)
		//add new server
		UserServers = append(UserServers, serverData)
	}
	return UserServers, nil
}
