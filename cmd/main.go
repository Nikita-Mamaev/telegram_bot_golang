package main

import (
	"github.com/Nikita-Mamaev/telegram_bot_golang/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1713704434:AAFT8OFoy4xYBYKARwuHSEP4pT9X8h7xCpk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	telegramBot.Start()

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
