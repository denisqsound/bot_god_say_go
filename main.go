package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var ChatId int64
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

	// TODO Получить ID чата
	ChatId = 123
	//fmt.Println(updates)
	// TODO Начать в цикле слать картинки
	for {
		//tgbotapi.NewPhotoUpload(ChatId,"https://assets.pinterest.com/ext/embed.html?id=68117013100642866" )
		msg := tgbotapi.NewMessage(ChatId, "qwe")
		bot.Send(msg)
		time.Sleep(10 * time.Second)
	}

	//for update := range updates {
	//	log.Printf("BOT UPDATE : %s", update.Message)
	//	if update.Message == nil { // ignore any non-Message Updates
	//		continue
	//		//tgbotapi.NewPhotoUpload(update.Message.Chat.ID,"https://assets.pinterest.com/ext/embed.html?id=68117013100642866")
	//		//tgbotapi.NewMessage(update.Message.Chat.ID, "qwe")
	//
	//	}
	//	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//	ChatId = update.Message.Chat.ID
	//
	//	//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	//	//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//	//msg.ReplyToMessageID = update.Message.MessageID
	//	//
	//	//bot.Send(msg)
	//}

}
