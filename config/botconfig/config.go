package botconfig

import (
	"log"
	"os"
)

var (
	BOTPORT = getBotPort()
)

func getBotPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("There is no BOTPORT!")
		port = "12345"
	}

	return port
}
