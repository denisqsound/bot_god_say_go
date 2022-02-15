package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// 401 картинка скачана 10 июля 2021 года

func main() {
	botToken, _ := os.LookupEnv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		log.Printf("BOT UPDATE : %s", update.Message)
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "/start" {
			for {
				fileName := fmt.Sprintf("bible_app/bible-%d.JPG", Random(1, 402))

				photoBytes, err := ioutil.ReadFile(fileName)
				if err != nil {
					log.Fatalln("Unable read file: ", err)
				}

				photoFileBytes := tgbotapi.FileBytes{
					Name:  "picture",
					Bytes: photoBytes,
				}

				_, _ = bot.Send(tgbotapi.NewPhotoUpload(update.Message.Chat.ID, photoFileBytes))
				_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Я родился"))

				if err != nil {
					log.Fatalln("Unable send photo: ", err)
				}
				// ToDo брать время из конфига
				time.Sleep(20 * time.Second)

			}

		}

	}

}

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min > max {
		return min
	} else {
		return rand.Intn(max-min) + min
	}
}
