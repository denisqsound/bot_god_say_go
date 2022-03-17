package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/denisqsound/bot_god_say_go/helpers"
	"github.com/denisqsound/bot_god_say_go/internal"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// 401 картинка скачана 10 июля 2021 года
// 516 уникальных картинок до 16 февраля 2022 года

func main() {

	botToken := internal.GetToken()
	timeSleep := internal.GetTime()
	internal.CheckFolderImages()

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Can't get update chan %v", err)
	}

	for update := range updates {

		log.Printf("BOT UPDATE : %v", update.Message)
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "/start" {
			for {
				pathImages, _ := os.LookupEnv("PICTURE_PATH")
				fileName := fmt.Sprintf("%v/bible-%d.JPG", pathImages, helpers.Random(1, 516))

				photoBytes, err := ioutil.ReadFile(fileName)
				if err != nil {
					log.Fatalln("Unable read file: ", err)
				}

				photoFileBytes := tgbotapi.FileBytes{
					Name:  "picture",
					Bytes: photoBytes,
				}

				//num := 1
				//text := fmt.Sprintf("Это %v", num)
				_, _ = bot.Send(tgbotapi.NewPhotoUpload(update.Message.Chat.ID, photoFileBytes))
				//_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text))

				if err != nil {
					log.Fatalln("Unable send photo: ", err)
				}

				time.Sleep(time.Duration(timeSleep) * time.Second)
			}
		}
	}
}
