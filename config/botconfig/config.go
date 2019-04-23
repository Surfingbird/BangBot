package botconfig

import (
	"log"
	"os"
)

var (
	BOTPORT     = getBotPort()
	VKVERSION   = "5.90"
	VKAPIADDR   = ""
	VKAPPID     = "6956956"
	SECUREKEY   = getSecureKey()
	ACCESSTOKEN = getAccessToken()
	MYVKID = "242753789"
)

func getBotPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("There is no BOTPORT!")
		port = "12345"
	}

	return port
}

func getAccessToken() string {
	token := os.Getenv("ACCESSTOKEN")
	if token == "" {
		log.Println("There is no ACCESSTOKEN!")
		token = "ACCESSTOKEN"
	}

	return token
}

func getSecureKey() string {
	key := os.Getenv("SECUREKEY")
	if key == "" {
		log.Println("There is no SECUREKEY!")
		key = "key"
	}

	return key
}
