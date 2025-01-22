package inituser

import "github.com/charmbracelet/huh"

func GetIpOfServer() string {
	var tmpIp string
	huh.NewInput().
		Title("Ip of server which will be delete?").
		Prompt("?").
		Value(&tmpIp).
		Run()
	return tmpIp
}
