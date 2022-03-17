package internal

import (
	"log"
	"os"
)

func CheckFolderImages() {
	pathImages := GetFolderImages()
	_, err := os.Stat(pathImages)
	if err != nil {
		log.Fatalf("Ошибка при поиске папки с картинками: %v", err)
	}
	if os.IsNotExist(err) {
		log.Fatalf("Папка %v НЕ существует \n", pathImages)
	}
	log.Printf("Папка %v найдена \n", pathImages)
}
