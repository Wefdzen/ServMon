package smtp

import (
	"fmt"
	"net/smtp"
)

// By default use smtp.gmail.com
// You will be message yourself
func SendMail(nameOfServer, status, loadAvg string) error {
	// Data for mail
	var (
		yourMail           = "yourEmail@gmail.com"
		passwordOfYourMail = "password"
		host               = "smtp.gmail.com"
	)
	// Set up authentication information.
	auth := smtp.PlainAuth("", yourMail, passwordOfYourMail, host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{yourMail}

	msg := []byte(fmt.Sprintf("To: %v\r\n"+
		"Subject: ServMon %v %v!\r\n"+
		"\r\n"+
		"Load Average last 5 min was %v.\r\n", yourMail, nameOfServer, status, loadAvg))
	//You can you yourself mail
	err := smtp.SendMail(fmt.Sprintf("%v:25", host), auth, yourMail, to, msg)
	if err != nil {
		return err
	}
	return nil
}
