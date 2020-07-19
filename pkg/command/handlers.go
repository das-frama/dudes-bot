package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"s32x.com/ovrstat/ovrstat"
)

const helpText = `
Список доступных команд:
/stop - Приостановить работу в этом чате;
/ping - Проверить доступность;
/schedule [день] - Работает ли сегодня Саша;
/overwatch [@участник] - Статистика игроков в Overwatch;
/cat - Призвать котика;
`

// Start runs the bot in chat on the first run or if it's was previously stopped.
func start(cfg commandConfig) (Result, error) {
	var result Result

	// Check if bot already started in the chat.
	_, created, err := cfg.Queryer.GetOrCreateChat(cfg.Message.Chat)
	if err != nil {
		return result, err
	}

	if created {
		result.Text = fmt.Sprintf("Я пробудился, как медведь от спячки зимней.\n%s", helpText)
	} else {
		result.Text = "Я уже запущен.\n/help - вызвать справку."
	}

	return result, nil
}

func stop(cfg commandConfig) (Result, error) {
	var result Result

	// Check if bot already started in the chat.
	err := cfg.Queryer.StopChat(cfg.Message.Chat.ID)
	if err != nil {
		return result, err
	}

	result.Text = "Я остановился.\nЧтобы запустить меня снова используйте /start."

	return result, nil
}

func ping(cfg commandConfig) (Result, error) {
	return Result{
		Text: "Да, и в самом деле я жив.",
	}, nil
}

// Schedule for Sasha.
func schedule(cfg commandConfig) (Result, error) {
	var result Result

	// Get first work day.
	workTime, err := time.Parse("02.01.2006", "05.06.2020")
	if err != nil {
		return result, err
	}

	// Params.
	params := cfg.Message.Params()

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

func call(cfg commandConfig) (Result, error) {
	var result Result

	result.Text = fmt.Sprintf(
		"Мне поступила команда, чтобы я всех призвал %s.",
		strings.Join(cfg.Message.Params(), " "),
	)

	return result, nil
}

func overwatch(cfg commandConfig) (Result, error) {
	var result Result

	// Default text.
	text := `Я пока не умею отображать статистику всех членов нашей команды.\n
	Пожалуйста, вызовите эту команду ещё раз, но с упоминанием кого-нибудь через @.`

	users := cfg.Message.Mentions()
	if len(users) > 0 {
		tags := make(map[string]string)

		// Get player tag.
		file, _ := ioutil.ReadFile("data/overwatch.json")
		_ = json.Unmarshal(file, &tags)
		if tag, ok := tags[users[0]]; ok {
			if stats, err := ovrstat.PCStats(tag); err != nil {
				text = "Ошибка"
			} else {
				text = ""
				for _, rating := range stats.Ratings {
					text += fmt.Sprintf("%s: %d\n", rating.Role, rating.Level)
				}
			}
		}
	}

	result.Text = text

	return result, nil
}

func cat(cfg commandConfig) (Result, error) {
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

func help(cfg commandConfig) (Result, error) {
	return Result{
		Text: helpText,
	}, nil
}
