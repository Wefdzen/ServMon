package main

import (
	"log"

	launcapp "github.com/Wefdzen/ServMon/internal/launcApp"
	"github.com/Wefdzen/ServMon/internal/website/router"
)

func main() {
	go launcapp.LaunchApp()
	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
