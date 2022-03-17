package internal

import (
	"log"
	"os"
	"strconv"
)

func GetTime() int64 {
	timer, ok := os.LookupEnv("WAIT_TIME")

	if !ok {
		log.Fatal("Can't get TIME from .env")
	}
	timeSleep, err := strconv.ParseInt(timer, 10, 64)
	if err != nil {
		log.Fatal("Can't convert TIME to INT")
	}

	log.Printf("Мы будем ждать %v секунд\n", timeSleep)
	return timeSleep

}
