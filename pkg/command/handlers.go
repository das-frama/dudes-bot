package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
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
/dog - Призвать пёселя;
/panda - Призвать панду;
`

const catURL = "https://cataas.com/cat"
const dogURL = "https://placedog.net/500"
const pandaURL = "https://loremflickr.com/500/500/panda"

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
	case "позапозавчера":
		currentTime = currentTime.Add(time.Hour * 24 * -3)
	case "позавчера":
		currentTime = currentTime.Add(time.Hour * 24 * -2)
	case "вчера":
		currentTime = currentTime.Add(time.Hour * 24 * -1)
	case "завтра":
		currentTime = currentTime.Add(time.Hour * 24)
	case "послезавтра":
		currentTime = currentTime.Add(time.Hour * 24 * 2)
	case "послепослезавтра":
		currentTime = currentTime.Add(time.Hour * 24 * 3)
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
	result := Result{
		PhotoURL: catURL,
	}

	// If query was made by Anna.
	if cfg.Message.From.Username == "unknow2n" {
		joke := "Когда Аня пытается понять в чём прикол"
		catURL := "https://cataas.com/cat/5ef820f05bc3fa0010444489"
		result.PhotoURL = fmt.Sprintf("%s/says/%s", catURL, joke)
	} else {
		// Retrieve joke.
		joke, err := cfg.Queryer.QueryRandomCatJoke()
		if err != nil {
			return result, err
		}
		if joke.Text != "" {
			result.PhotoURL = fmt.Sprintf("%s/says/%s", result.PhotoURL, joke.Text)
			fmt.Println(result.PhotoURL)
		}
		// Random seed number for new image to appear.
		rand := rand.Int()
		result.PhotoURL = fmt.Sprintf("%s?seed=%d", result.PhotoURL, rand)
	}

	return result, nil
}

func dog(cfg commandConfig) (Result, error) {
	url := fmt.Sprintf("%s?random&seed=%d", dogURL, rand.Int())

	return Result{
		PhotoURL: url,
	}, nil
}

func panda(cfg commandConfig) (Result, error) {
	url := fmt.Sprintf("%s?random=%d", pandaURL, rand.Int())

	return Result{
		PhotoURL: url,
	}, nil
}

func help(cfg commandConfig) (Result, error) {
	return Result{
		Text: helpText,
	}, nil
}
