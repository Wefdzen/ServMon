package workwithservers

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

// SendCommandToServer ip-ip of server, pasw - pasw of server, command - command to request
// portSSH - if empty like("") will be by default 22 if not empty will be custom port
// return a output of your command.
func SendCommandToServer(ip, userRole, password, command, portSSH string) (string, error) {
	// data for connect to server
	server := ip     // ip of server
	user := userRole // name of user
	port := "22"
	if portSSH != "" {
		port = portSSH
	}

	// Config for SSH
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Важно: для простоты, но не рекомендуется в продакшн
	}

	// Connect to server
	addr := fmt.Sprintf("%s:%s", server, port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// open new session
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// execute a command
	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}
	//return output of command
	return string(output), nil
}
