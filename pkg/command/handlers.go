package command

import (
	"das-frama/dudes-bot/pkg/bot"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"s32x.com/ovrstat/ovrstat"
)

const firstWorkDay = "05.06.2020"

func start(update bot.Update) CommandResponse {
	fileName := "data/chats.json"
	var chats []*bot.Chat

	// If file exists.
	if _, err := os.Stat(fileName); err == nil {
		file, _ := ioutil.ReadFile(fileName)
		json.Unmarshal(file, &chats)
	}

	// Check if chat is in slice.
	if isExist := inChats(chats, update.Message.Chat.ID); !isExist {
		chats = append(chats, update.Message.Chat)
	}

	// Store chat.
	file, _ := json.MarshalIndent(chats, "", "  ")
	_ = ioutil.WriteFile("data/chats.json", file, 0644)

	// Reply.
	return CommandResponse{
		Reply: "А я уже запущен",
	}
}

func ping(update bot.Update) CommandResponse {
	return CommandResponse{
		Reply: "pong",
	}
}

func schedule(update bot.Update) CommandResponse {
	// workTime, err := time.Parse("02.01.2006", firstWorkDay)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// var verb string
	// if msg == "" {
	// 	verb = "Сегодня"
	// } else {
	// 	verb = msg
	// }

	// currentTime := time.Now()
	// switch msg {
	// case "завтра":
	// 	currentTime = currentTime.Add(time.Hour * 24)
	// case "послезавтра":
	// 	currentTime = currentTime.Add(time.Hour * 48)
	// }

	// hours := currentTime.Sub(workTime).Hours()
	// days := int(hours / 24)

	// var reply string
	// if days%4 < 2 {
	// 	reply = "%s Саша трудится в поте лица!"
	// } else {
	// 	reply = "%s Саша отдыхает! 😊😊😊"
	// }

	// return CommandResponse{
	// 	Reply: fmt.Sprintf(reply, strings.Title(verb)),
	// }
	return CommandResponse{
		Reply: "schedule",
	}
}

func call(update bot.Update) CommandResponse {
	// reply := fmt.Sprintf("Мне поступила команда, чтобы я всех призвал %s.", msg)
	reply := fmt.Sprintf("Мне поступила команда, чтобы я всех призвал!")
	return CommandResponse{
		Reply: reply,
	}
}

func overwatch(update bot.Update) CommandResponse {
	reply := "Я пока не умею отображать статистику всех членов нашей славной команды. Пожалуйста, вызовите эту команду ещё раз, но с упомянием кого-нибудь через `@`."
	users := update.Message.Mentions()
	if len(users) > 0 {
		tags := make(map[string]string)

		// Get player tag.
		file, _ := ioutil.ReadFile("data/overwatch.json")
		_ = json.Unmarshal(file, &tags)
		if tag, ok := tags[users[0]]; ok {
			if stats, err := ovrstat.PCStats(tag); err != nil {
				reply = "Ошибка"
			} else {
				reply = ""
				for _, rating := range stats.Ratings {
					reply += fmt.Sprintf("%s: %d\n", rating.Role, rating.Level)
				}
			}
		}

	}

	return CommandResponse{
		Reply: reply,
	}
}
