package main

import (
	"log"

	"github.com/AlexBragin1/golang_telegram_bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("")

	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	telegramBot := telegram.NewBot(bot)

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
func initLoad() {
	//if err:=godoenv.Load();err!=nil{

	// }
}
