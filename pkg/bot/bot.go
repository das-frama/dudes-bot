package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// BaseURL
const BaseURL = "https://api.telegram.org/bot%s/%s"

type Bot struct {
	Token  string
	Buffer int

	shutdownChannel chan interface{}
}

// New creates a new bot.
func New(token string) *Bot {
	return &Bot{
		Token:  token,
		Buffer: 100,

		shutdownChannel: make(chan interface{}),
	}
}

// GetUpdatesChan starts and returns a channel for getting updates.
func (bot *Bot) GetUpdatesChan(config UpdateConfig) (UpdatesChannel, error) {
	ch := make(chan Update, bot.Buffer)

	go func() {
		for {
			select {
			case <-bot.shutdownChannel:
				close(ch)
				return
			default:
			}

			updates, err := bot.GetUpdates(config)
			if err != nil {
				log.Println(err)
				log.Println("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)
				continue
			}

			for _, update := range updates {
				if update.UpdateID >= config.Offset {
					config.Offset = update.UpdateID + 1
					ch <- update
				}
			}
		}
	}()

	return ch, nil
}

// GetUpdates eceive incoming updates using long polling (wiki). An Array of Update objects is returned.
func (bot *Bot) GetUpdates(config UpdateConfig) ([]Update, error) {
	v := url.Values{}

	if config.Offset != 0 {
		v.Add("offset", strconv.Itoa(config.Offset))
	}
	if config.Limit > 0 {
		v.Add("limit", strconv.Itoa(config.Limit))
	}
	if config.Timeout > 0 {
		v.Add("timeout", strconv.Itoa(config.Timeout))
	}

	response, err := bot.request("getUpdates", v)
	if err != nil {
		return []Update{}, err
	}

	var updates []Update
	json.Unmarshal(response.Result, &updates)

	return updates, nil
}

// SendMessage send text messages. On success, the sent Message is returned.
func (bot *Bot) SendMessage(config SendMessageConfig) (Message, error) {
	v, _ := config.values()

	response, err := bot.request("sendMessage", v)
	if err != nil {
		return Message{}, nil
	}

	var message Message
	json.Unmarshal(response.Result, &message)

	return message, nil
}

func (bot *Bot) request(method string, params url.Values) (Response, error) {
	url := fmt.Sprintf(BaseURL, bot.Token, method)

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return Response{}, err
	}

	if !response.Ok {
		return response, fmt.Errorf("not ok")
	}

	return response, nil
}
