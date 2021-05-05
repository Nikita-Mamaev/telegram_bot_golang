package main

import (
	"github.com/Nikita-Mamaev/telegram_bot_golang/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient(pocketToken)
	if err != nil {
		log.Fatal(err)
	}
	telegramBot := telegram.NewBot(bot, pocketClient, "http://localhost")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
