package bot

// UpdateConfig stands for config for Update method.
type UpdateConfig struct {
	Offset  int
	Limit   int
	Timeout int
}

type SendMessageConfig struct {
	ChatID                int             `json:"chat_id"`
	Text                  string          `json:"text"`
	ParseMode             string          `json:"parse_mode"`
	DisableWebPagePreview bool            `json:"disable_web_page_preview"`
	DisableNotification   bool            `json:"disable_notification"`
	ReplyToMessageID      int             `json:"reply_to_message_id"`
	Entities              []MessageEntity `json:"entities"`
}

type SendPhotoConfig struct {
	ChatID              int
	Photo               string
	ParseMode           string
	Caption             string
	DisableNotification bool
	ReplyToMessageID    int
}

// func (config *SendMessageConfig) values() (url.Values, error) {
// 	v := url.Values{}
// 	v.Add("chat_id", strconv.Itoa(config.ChatID))
// 	v.Add("text", config.Text)
// 	v.Add("parse_mode", config.ParseMode)
// 	v.Add("disable_notification", strconv.FormatBool(config.DisableNotification))
// 	v.Add("disable_web_page_preview", strconv.FormatBool(config.DisableWebPagePreview))
// 	v.Add("reply_to_message_id", strconv.Itoa(config.ReplyToMessageID))
// 	return v, nil
// }

// func (config *SendPhotoConfig) values() (url.Values, error) {
// 	v := url.Values{}
// 	v.Add("chat_id", strconv.Itoa(config.ChatID))
// 	v.Add("photo", config.Photo)
// 	v.Add("caption", config.Caption)
// 	v.Add("parse_mode", config.ParseMode)
// 	v.Add("disable_notification", strconv.FormatBool(config.DisableNotification))
// 	v.Add("reply_to_message_id", strconv.Itoa(config.ReplyToMessageID))
// 	return v, nil
// }
