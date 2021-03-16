package bot

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// BaseURL is base url for telegram api.
const BaseURL = "https://api.telegram.org/bot%s/%s"

// Bot struct represents a single bot.
type Bot struct {
	Token  string
	Buffer int

	db              *sql.DB
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
	jsonStr, err := json.Marshal(config)
	if err != nil {
		return []Update{}, err
	}

	response, err := bot.request("getUpdates", jsonStr)
	if err != nil {
		return []Update{}, err
	}

	var updates []Update
	json.Unmarshal(response.Result, &updates)

	return updates, nil
}

// SendMessage send text messages. On success, the sent Message is returned.
func (bot *Bot) SendMessage(config SendMessageConfig) (Message, error) {
	// Prepare json string.
	jsonStr, err := json.Marshal(config)
	if err != nil {
		return Message{}, err
	}
	fmt.Println(bytes.NewBuffer(jsonStr))

	// Send request to telegram.
	response, err := bot.request("sendMessage", jsonStr)
	if err != nil {
		return Message{}, err
	}

	// Unmarshal response.
	var message Message
	json.Unmarshal(response.Result, &message)

	return message, nil
}

// SendPhoto send text messages. On success, the sent Message is returned.
func (bot *Bot) SendPhoto(config SendPhotoConfig) (Message, error) {
	jsonStr, err := json.Marshal(config)
	if err != nil {
		return Message{}, err
	}

	response, err := bot.request("sendPhoto", jsonStr)
	if err != nil {
		return Message{}, nil
	}

	var message Message
	json.Unmarshal(response.Result, &message)

	return message, nil
}

func (bot *Bot) request(method string, jsonBody []byte) (Response, error) {
	url := fmt.Sprintf(BaseURL, bot.Token, method)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
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
		return response, fmt.Errorf(response.Description)
	}

	return response, nil
}
