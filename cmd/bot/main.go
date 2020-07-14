package main

import (
	"das-frama/dudes-bot/pkg/bot"
	"das-frama/dudes-bot/pkg/command"
	"das-frama/dudes-bot/pkg/config"
	"das-frama/dudes-bot/pkg/sqlite"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Get bot api token from env.
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("$BOT_TOKEN must be set.")
	}

	// Load config.
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Read config: ", config)

	// Open DB.
	db, err := sqlite.LoadDB(config.DB.Path)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Init DB.
	err = sqlite.Migrate(db, config.DB.Init)
	if err != nil {
		log.Fatalln(err)
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

		// Check if message is command.
		if update.Message.IsCommand() {
			cmd := update.Message.Command()
			params := update.Message.Params()
			result, err := command.Process(cmd, params)
			if err != nil {
				log.Println(err)
			}
			// Send message.
			_, err = tgBot.SendMessage(bot.SendMessageConfig{
				ChatID: update.Message.Chat.ID,
				Text:   result.Text,
			})
			if err != nil {
				log.Println(err)
			}
		}

	}

	log.Println("Shutting down...")
}
