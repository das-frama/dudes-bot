package command

import (
	"das-frama/dudes-bot/pkg/bot"
	"fmt"
	"strings"
	"time"
	// "s32x.com/ovrstat"
)

// HandlerConfig
type HandlerConfig struct {
	Update     *bot.Update
	RowQueryer RowQueryer
}

// Start runs the bot in chat on the first run or if it's was previously stopped.
func start(params []string) (Result, error) {
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
	return Result{
		Text: "start",
	}, nil
}

func stop(params []string) (Result, error) {
	return Result{
		Text: "stop",
	}, nil
}

func ping(params []string) (Result, error) {
	return Result{
		Text: "pong",
	}, nil
}

// Schedule for Sasha.
func schedule(params []string) (Result, error) {
	var result Result

	// Get first work day.
	workTime, err := time.Parse("02.01.2006", "05.06.2020")
	if err != nil {
		return result, err
	}

	// Get word.
	word := "сегодня"
	if len(params) > 0 && params[0] != "" {
		word = params[0]
	}

	// Find out word offset.
	currentTime := time.Now()
	switch word {
	case "завтра":
		currentTime = currentTime.Add(time.Hour * 24)
	case "послезавтра":
		currentTime = currentTime.Add(time.Hour * 48)
	case "послепослезавтра":
		currentTime = currentTime.Add(time.Hour * 72)
	}

	// Calculate days
	days := int(currentTime.Sub(workTime).Hours() / 24)

	var text string
	if days%4 < 2 {
		text = "%s Саша трудится в поте лица!"
	} else {
		text = "%s Саша отдыхает! 😊😊😊"
	}
	result.Text = fmt.Sprintf(text, strings.Title(word))

	return result, nil
}

func call(params []string) (Result, error) {
	// Text := fmt.Sprintf("Мне поступила команда, чтобы я всех призвал %s.", msg)
	// Text := fmt.Sprintf("Мне поступила команда, чтобы я всех призвал!")
	return Result{
		// Text: Text,
		Text: "call",
	}, nil
}

func overwatch(params []string) (Result, error) {
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

	return Result{
		// Text: Text,
		Text: "overwatch",
	}, nil
}

func cat(params []string) (Result, error) {
	// 	resp, err := http.Get("https://cataas.com/cat")

	// 	buffer := &bytes.Buffer
	// 	w := multipart.NewWriter(buffer)
	// 	part := w.create
	// 	if err != nil {
	// 		log.Print(err)
	// 	}

	return Result{
		Text: "cat",
	}, nil
}
