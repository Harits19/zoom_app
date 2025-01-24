package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ClientId, ClientSecret, ApplicationId string

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ClientId = os.Getenv("ClientId")
	ClientSecret = os.Getenv("ClientSecret")
	ApplicationId = os.Getenv("ApplicationId")

	fmt.Println("Check ENV")
	fmt.Println("ClientId", ClientId)
	fmt.Println("ClientSecret", ClientSecret)
	fmt.Println("ApplicationId", ApplicationId)
}
