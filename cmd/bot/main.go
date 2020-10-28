package main

import (
	"das-frama/dudes-bot/pkg/bot"
	"das-frama/dudes-bot/pkg/command"
	"das-frama/dudes-bot/pkg/config"
	"das-frama/dudes-bot/pkg/sqlite"
	"database/sql"
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

	// Config.
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	// Setup dependencies.
	// DB Conn.
	conn, err := sql.Open("sqlite3", cfg.DB.Path)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// DB struct.
	db := sqlite.New(conn)
	// Init schema.
	err = db.InitSchema(cfg.DB.Init)
	if err != nil {
		log.Fatalln(err)
	}

	// Create telegram bot object.
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
		// if ok, err := db.IsChatActive(update.Message.Chat.ID); !ok && err == nil {
		// continue
		// }

		// Log incoming message.
		log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)

		// Check if message is command.
		if update.Message.IsCommand() {
			cmd := update.Message.Command()
			result, err := command.Process(cmd, update.Message, db, cfg)
			if err != nil {
				result.Text = command.UcFirst(err.Error()) + "."
				log.Println(err)
			}
			// Send message.
			if result.Text != "" {
				_, err = tgBot.SendMessage(bot.SendMessageConfig{
					ChatID: update.Message.Chat.ID,
					Text:   result.Text,
				})
				if err != nil {
					log.Println(err)
				}
			} else if result.PhotoURL != "" {
				_, err = tgBot.SendPhoto(bot.SendPhotoConfig{
					ChatID: update.Message.Chat.ID,
					Photo:  result.PhotoURL,
				})
				if err != nil {
					log.Println(err)
				}
			}
		}

	}

	log.Println("Shutting down...")
}
