package command

import "das-frama/dudes-bot/pkg/bot"

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
