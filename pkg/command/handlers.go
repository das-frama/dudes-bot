package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
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
const memeURL = "https://meme-api.herokuapp.com/gimme"

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
	workDate, err := time.Parse("2006-01-02 MST", "2020-10-27 MSK")
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
	var nextDate time.Time
	now := time.Now()
	isDate := false
	isNextYear := false
	switch word {
	case "сегодня":
		nextDate = now
	case "позапозавчера":
		nextDate = now.Add(time.Hour * 24 * -3)
	case "позавчера":
		nextDate = now.Add(time.Hour * 24 * -2)
	case "вчера":
		nextDate = now.Add(time.Hour * 24 * -1)
	case "завтра":
		nextDate = now.Add(time.Hour * 24)
	case "послезавтра":
		nextDate = now.Add(time.Hour * 24 * 2)
	case "послепослезавтра":
		nextDate = now.Add(time.Hour * 24 * 3)
	default:
		word = fmt.Sprintf("%s.%d %s MSK", word, now.Year(), now.Format("03:04:05.999999999"))
		if nextDate, err = time.Parse("02.01.2006 03:04:05.999999999 MST", word); err != nil {
			return result, ErrWrongDateFormat
		}
		if nextDate.Before(now) {
			nextDate = nextDate.AddDate(1, 0, 0)
			isNextYear = true
		}
		isDate = true
	}

	// Calculate days
	days := int(nextDate.Sub(workDate).Hours() / 24)

	// Found out what text should be displayed.
	var text string
	if days%4 < 2 {
		text = "Саша трудится в поте лица!"
	} else if nextDate.Day() == 10 && nextDate.Month() == time.November {
		text = "Саша отдыхает, но встречать Андрея не поедет! 😢😢😢"
	} else {
		text = "Саша отдыхает! 😊😊😊"
	}

	// Fill up the result struct.
	if isDate {
		if isNextYear {
			result.Text = fmt.Sprintf("%d %s %d %s", nextDate.Day(), months[nextDate.Month()-1], nextDate.Year(), text)
		} else {
			result.Text = fmt.Sprintf("%d %s %s", nextDate.Day(), months[nextDate.Month()-1], text)
		}
	} else {
		result.Text = fmt.Sprintf("%s %s", strings.Title(word), text)
	}

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

func meme(cfg commandConfig) (Result, error) {
	var result Result

	resp, err := http.Get(memeURL)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	var meme MemeResponse
	json.NewDecoder(resp.Body).Decode(&meme)
	result.PhotoURL = meme.URL

	return result, nil
}

func help(cfg commandConfig) (Result, error) {
	return Result{
		Text: helpText,
	}, nil
}
