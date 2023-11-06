package main

import (
	"flag"
	tg "getting-started-tgBotAPI/clients/telegram"
	"getting-started-tgBotAPI/consumer/eventConsumer"
	"getting-started-tgBotAPI/events/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	/*
		fmt.Println("Bot is online!")

		os.Setenv("TELEGRAM_APITOKEN", "123456789")

		bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
		if err != nil {
			panic(err)
		}

		bot.Debug = true

		updateConfig := tgbotapi.NewUpdate(0)

		updateConfig.Timeout = 30

		updates := bot.GetUpdatesChan(updateConfig)

		for update := range updates {
			if update.Message == nil {
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	*/

	t := mustToken()

	tgClient := tg.New(tgBotHost, t)

	eventProcessor := telegram.New(tgClient)

	log.Print("service started")

	consumer := eventConsumer.New(eventProcessor, eventProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stropped")
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
