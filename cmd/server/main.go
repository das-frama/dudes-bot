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

	// Get updates channel.
	updates, err := tgBot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatalln("error")
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)
		var reply string
		switch update.Message.Command() {
		case "start":
			reply = "А я уже запущен."
		case "joke":
			reply = "Если ваш крем от морщин действует, то почему у вас до сих пор есть отпечатки пальцев?"
		case "judge":
			reply = "Хороший выбор! Однако я ещё не научился быть как Дредд, поэтому пускай будет прав тот, кто первым меня об этом попросил."
		case "bk":
			reply = "Хм... Дайте подумать... Простите, я не знаю, что такое 'БК'. Если это какой-то ресторан быстрого питания, то боюсь, я огорчу вас своим ответом."
		}
		if reply != "" {
			tgBot.SendMessage(bot.SendMessageConfig{
				ChatID: update.Message.Chat.ID,
				Text:   reply,
			})
		}
	}
}
