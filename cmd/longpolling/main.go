package main

import (
	"das-frama/dudes-bot/pkg/bot"
	"das-frama/dudes-bot/pkg/command"
	"log"
	"os"
)

func main() {
	// Get bot api token.
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("$BOT_TOKEN must be set.")
	}

	// Create telegram bot instance.
	tgBot := bot.New(token)

	// Get updates channel.
	updates, err := tgBot.GetUpdatesChan(bot.UpdateConfig{
		Offset:  0,
		Limit:   0,
		Timeout: 60,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Get through channels.
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Log incoming message.
		log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)

		// If message is command then proccess it.
		if update.Message.IsCommand() {
			cmd := update.Message.Command()
			reply, err := command.Process(cmd, update)
			if err != nil {
				log.Println(err.Error())
			}
			if reply != "" {
				// Send message.
				tgBot.SendMessage(bot.SendMessageConfig{
					ChatID: update.Message.Chat.ID,
					Text:   reply,
				})
			}
		}
	}

	log.Println("Shutting down...")
}
