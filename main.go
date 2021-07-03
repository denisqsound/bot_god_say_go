package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/levigross/grequests"
    "github.com/robfig/cron/v3"
	"log"
	"os"
	"strconv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Точка входа программы
func main() {
	botToken, _ := os.LookupEnv("TOKEN")
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	offset := 0

	updates, err := getUpdates(botUrl, offset)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	c := cron.New()
	c.AddFunc("@every second", Respond(botUrl, updates)
	{
		fmt.Printf("Every second ")
	})
	c.Start()

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
func getUpdates(botUrl string, offset int) ([]Update, error) {
	ro := &grequests.RequestOptions{Params: map[string]string{"offset": strconv.Itoa(offset)}}
	res, err := grequests.Get(botUrl+"/getUpdates", ro)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	resp := ResResponse{}
	err = res.JSON(&resp)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	return resp.Result, nil

}

// Отвечает на обновление

func SendMessage(botUrl string, update Update) {
	botMessage := BotMessage{}
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Photo = "https://www.pinterest.com/pin/68117013100642866/"

	ro := &grequests.RequestOptions{
		Headers: map[string]string{"Content-Type": "application/json"},
		JSON:    &botMessage,
	}
	_, err := grequests.Post(botUrl+"/sendPhoto", ro)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	return

}
