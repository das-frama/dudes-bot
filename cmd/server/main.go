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
		reply := "И хоть это странно, что я реагирую на каждую фразу тем же сообщением, я нахожу в этом красоту и равновесие."
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)
		switch update.Message.Command() {
		case "start":
			reply = "А я уже запущен."
		case "joke":
			reply = "Если ваш крем от морщин действует, то почему у вас до сих пор есть отпечатки пальцев?"
		}
		tgBot.SendMessage(bot.SendMessageConfig{
			ChatID: update.Message.Chat.ID,
			Text:   reply,
		})
	}

}
