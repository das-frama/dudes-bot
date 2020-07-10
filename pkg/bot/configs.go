package bot

import (
	"net/url"
	"strconv"
)

type UpdateConfig struct {
	Offset  int
	Limit   int
	Timeout int
}

type SendMessageConfig struct {
	ChatID                int
	Text                  string
	ParseMode             string
	DisableWebPagePreview bool
	DisableNotification   bool
	ReplyToMessageID      int
}

type SendPhotoConfig struct {
	ChatID              int
	Photo               string
	ParseMode           string
	Caption             string
	DisableNotification bool
	ReplyToMessageID    int
}

func (config *SendMessageConfig) values() (url.Values, error) {
	v := url.Values{}
	v.Add("chat_id", strconv.Itoa(config.ChatID))
	v.Add("text", config.Text)
	v.Add("parse_mode", config.ParseMode)
	v.Add("disable_notification", strconv.FormatBool(config.DisableNotification))
	v.Add("disable_web_page_preview", strconv.FormatBool(config.DisableWebPagePreview))
	v.Add("reply_to_message_id", strconv.Itoa(config.ReplyToMessageID))
	return v, nil
}

func (config *SendPhotoConfig) values() (url.Values, error) {
	v := url.Values{}
	v.Add("chat_id", strconv.Itoa(config.ChatID))
	v.Add("photo", config.Photo)
	v.Add("caption", config.Caption)
	v.Add("parse_mode", config.ParseMode)
	v.Add("disable_notification", strconv.FormatBool(config.DisableNotification))
	v.Add("reply_to_message_id", strconv.Itoa(config.ReplyToMessageID))
	return v, nil
}
