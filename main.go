package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Точка входа программы
func main() {
	botToken, _ := os.LookupEnv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

	//////////////////////////////////////////////////

	//botToken, _ := os.LookupEnv("TOKEN")
	//botApi := "https://api.telegram.org/bot"
	//botUrl := botApi + botToken
	//offset := 0

	//updates, err := getUpdates(botUrl, offset)
	//if err != nil {
	//	log.Fatalln("Unable to make request: ", err)
	//}

	//c := cron.New()
	//c.AddFunc("@every 10s", func() { SendMessage(botUrl, updates) })
	//c.Start()

	//for {
	//	updates, err := getUpdates(botUrl, offset)
	//	if err != nil {
	//		log.Fatalln("Unable to make request: ", err)
	//	}
	//
	//	for _, update := range updates {
	//		time.Sleep(time.Second * 5)
	//		Respond(botUrl, update)
	//		offset = update.UpdateId + 1
	//
	//	}
	//	fmt.Println(updates)
	//}

}

// Запрос обновлений
//func getUpdates(botUrl string, offset int) ([]Update, error) {
//	ro := &grequests.RequestOptions{Params: map[string]string{"offset": strconv.Itoa(offset)}}
//	res, err := grequests.Get(botUrl+"/getUpdates", ro)
//	if err != nil {
//		log.Fatalln("Unable to make request: ", err)
//	}
//
//	resp := ResResponse{}
//	err = res.JSON(&resp)
//	if err != nil {
//		log.Fatalln("Unable to make request: ", err)
//	}
//
//	return resp.Result, nil
//
//}

// Отвечает на обновление
//func SendMessage(botUrl string, update Update) {
//	botMessage := BotMessage{}
//	botMessage.ChatId = update.Message.Chat.ChatId
//	botMessage.Photo = "https://www.pinterest.com/pin/68117013100642866/"
//
//	ro := &grequests.RequestOptions{
//		Headers: map[string]string{"Content-Type": "application/json"},
//		JSON:    &botMessage,
//	}
//	_, err := grequests.Post(botUrl+"/sendPhoto", ro)
//	if err != nil {
//		log.Fatalln("Unable to make request: ", err)
//	}
//
//	return
//
//}
