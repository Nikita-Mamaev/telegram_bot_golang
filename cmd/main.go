package main

import (
	"github.com/Nikita-Mamaev/telegram_bot_golang/pkg/repository"
	"github.com/Nikita-Mamaev/telegram_bot_golang/pkg/repository/boltdb"
	"github.com/Nikita-Mamaev/telegram_bot_golang/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

const ()

func main() {
	bot, err := tgbotapi.NewBotAPI("1713704434:AAFT8OFoy4xYBYKARwuHSEP4pT9X8h7xCpk")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	telegramBot.Start()

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}

		_, err := tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
