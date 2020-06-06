package main

import (
	"das-frama/dudes-bot/pkg/bot"
	"log"
	"os"
)

func main() {
	token := os.Getenv("BOT_TOKEN")

	if token == "" {
		log.Fatalln("$BOT_TOKEN must be set.")
	}

	tgBot := bot.New(token)
	updateConfig := bot.UpdateConfig{
		Offset:  0,
		Limit:   0,
		Timeout: 60,
	}

	updates, err := tgBot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatalln("error")
	}

	for update := range updates {
		// reply := "Не знаю, что вам сказать на это."
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)
	}

}
