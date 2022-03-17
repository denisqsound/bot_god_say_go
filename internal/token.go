package internal

import (
	"log"
	"os"
)

func GetToken() string {
	botToken, ok := os.LookupEnv("TOKEN")
	if !ok {
		log.Fatal("Can't get TOKEN from .env")
	}
	return botToken
}
