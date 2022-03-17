package internal

import (
	"log"
	"os"
)

func GetFolderImages() string {
	pathImages, ok := os.LookupEnv("PICTURE_PATH")
	if !ok {
		log.Fatal("Can't get PICTURE_PATH from .env")
	}
	return pathImages

}
