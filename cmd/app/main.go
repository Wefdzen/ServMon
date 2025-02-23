package main

import (
	"log"

	launchapp "github.com/Wefdzen/ServMon/internal/launchApp"
	"github.com/Wefdzen/ServMon/internal/website/router"
)

func main() {
	go launchapp.LaunchApp()
	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
