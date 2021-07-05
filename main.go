package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
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

var IMG = map[int]string{
	1:  "https://disk.yandex.ru/i/MYT2df5Iyyzg1g",
	2:  "https://disk.yandex.ru/i/RAHfOqdMJ-5bJQ",
	3:  "https://disk.yandex.ru/i/Agd2Fw4yTa4m6A",
	4:  "https://disk.yandex.ru/i/m6iPlUjpPfaMFg",
	5:  "https://disk.yandex.ru/i/QU-j_ziPBl41Yw",
	6:  "https://disk.yandex.ru/i/U1z4iO84HrT_ZA",
	7:  "https://disk.yandex.ru/i/7NUlGmFqkWU9Yg",
	8:  "https://disk.yandex.ru/i/m8Dr6_duw_Ecug",
	9:  "https://disk.yandex.ru/i/ABWTvj9aQR6p6A",
	10: "https://disk.yandex.ru/i/68FxAdPv9Yatmw",
	11: "https://disk.yandex.ru/i/VLv5MzsbZC9raA",
	12: "https://disk.yandex.ru/i/oUENMBUeQhhBiA",
	13: "https://disk.yandex.ru/i/C7z9tjPdymgzzA",
	14: "https://disk.yandex.ru/i/2CA3NuNzJ2fg5w",
	15: "https://disk.yandex.ru/i/KaqkPzLXd4J1nA",
	16: "https://disk.yandex.ru/i/FDagx_bvm2ciRQ",
	17: "https://disk.yandex.ru/i/jgG5lFMZCrmsXg",
	18: "https://disk.yandex.ru/i/6fR444jtFfkOYg",
	19: "https://disk.yandex.ru/i/DeybnWhYkNk05A",
	20: "https://disk.yandex.ru/i/N6jKIbd969MdcA",
}

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
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, RandomImage())
				time.Sleep(5 * time.Second)
				bot.Send(msg)
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

func RandomImage() string {
	return IMG[Random(1, 20)]

}
