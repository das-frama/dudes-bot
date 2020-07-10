package command

import (
	"das-frama/dudes-bot/pkg/bot"
	// "s32x.com/ovrstat"
)

// Start runs the bot in chat on the first run or if it's was previously stopped.
func start(update bot.Update) Response {
	// fileName := "data/chats.json"
	// var chats []*bot.Chat

	// // If file exists.
	// if _, err := os.Stat(fileName); err == nil {
	// 	file, _ := ioutil.ReadFile(fileName)
	// 	json.Unmarshal(file, &chats)
	// }

	// // Check if chat is in slice.
	// if isExist := inChats(chats, update.Message.Chat.ID); !isExist {
	// 	chats = append(chats, update.Message.Chat)
	// }

	// // Store chat.
	// file, _ := json.MarshalIndent(chats, "", "  ")
	// _ = ioutil.WriteFile("data/chats.json", file, 0644)

	// Text.
	return Response{
		Text: "start",
	}
}

func stop(update bot.Update) Response {
	return Response{
		Text: "stop",
	}
}

func ping(update bot.Update) Response {
	return Response{
		Text: "pong",
	}
}

func schedule(update bot.Update) Response {
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
	// case "послепослезавтра":
	// 	currentTime = currentTime.Add(time.Hour * 72)
	// }

	// hours := currentTime.Sub(workTime).Hours()
	// days := int(hours / 24)

	// var Text string
	// if days%4 < 2 {
	// 	Text = "%s Саша трудится в поте лица!"
	// } else {
	// 	Text = "%s Саша отдыхает! 😊😊😊"
	// }

	return Response{
		// Text: fmt.Sprintf(Text, strings.Title(verb)),
		Text: "schedule",
	}
}

func call(update bot.Update) Response {
	// Text := fmt.Sprintf("Мне поступила команда, чтобы я всех призвал %s.", msg)
	// Text := fmt.Sprintf("Мне поступила команда, чтобы я всех призвал!")
	return Response{
		// Text: Text,
		Text: "call",
	}
}

func overwatch(update bot.Update) Response {
	// Text := "Я пока не умею отображать статистику всех членов нашей славной команды. Пожалуйста, вызовите эту команду ещё раз, но с упомянием кого-нибудь через `@`."
	// users := update.Message.Mentions()
	// if len(users) > 0 {
	// 	tags := make(map[string]string)

	// 	// Get player tag.
	// 	file, _ := ioutil.ReadFile("data/overwatch.json")
	// 	_ = json.Unmarshal(file, &tags)
	// 	if tag, ok := tags[users[0]]; ok {
	// 		if stats, err := ovrstat.PCStats(tag); err != nil {
	// 			Text = "Ошибка"
	// 		} else {
	// 			Text = ""
	// 			for _, rating := range stats.Ratings {
	// 				Text += fmt.Sprintf("%s: %d\n", rating.Role, rating.Level)
	// 			}
	// 		}
	// 	}

	// }

	return Response{
		// Text: Text,
		Text: "overwatch",
	}
}

func cat(update bot.Update) Response {
	// 	resp, err := http.Get("https://cataas.com/cat")

	// 	buffer := &bytes.Buffer
	// 	w := multipart.NewWriter(buffer)
	// 	part := w.create
	// 	if err != nil {
	// 		log.Print(err)
	// 	}

	return Response{
		Text: "cat",
	}
}
