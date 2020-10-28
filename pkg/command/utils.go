package command

import (
	"das-frama/dudes-bot/pkg/bot"
	"unicode"
)

// MemeResponse is a struct for response for https://meme-api.herokuapp.com/gimme
type MemeResponse struct {
	PostLink  string
	Subreddit string
	Title     string
	URL       string
	Nsfw      bool
	Spoiler   bool
}

func inChats(chats []*bot.Chat, id int) bool {
	for _, chat := range chats {
		if chat.ID == id {
			return true
		}
	}
	return false
}

func isMention(entities []bot.MessageEntity) bool {
	for _, entity := range entities {
		if entity.Type == "mention" {
			return true
		}
	}

	return false
}

// UcFirst make first letter uppercase.
func UcFirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}

	return ""
}
