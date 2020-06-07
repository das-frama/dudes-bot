package bot

import (
	"encoding/json"
	"strings"
)

type Response struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}

// UpdatesChannel is the channel for getting updates.
type UpdatesChannel <-chan Update

// Update represents an incoming update.
type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	Poll               *Poll               `json:"poll"`
	PollAnswer         *PollAnswer         `json:"poll_answer"`
}

type InlineQuery struct {
}

type ChosenInlineResult struct {
}

type CallbackQuery struct {
}

type Poll struct {
}

type PollAnswer struct {
}

// User represents a Telegram user or bot.
type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

// Chat represents a chat.
type Chat struct {
	ID               int              `json:"id"`
	Type             string           `json:"type"`
	Title            string           `json:"title"`
	Username         string           `json:"username"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Photo            *ChatPhoto       `json:"photo"`
	Description      string           `json:"description"`
	InviteLink       string           `json:"invite_link"`
	PinnedMessage    *Message         `json:"pinned_message"`
	Permissions      *ChatPermissions `json:"permissions"`
	SlowModeDelay    int              `json:"slow_mode_delay"`
	StickerSetName   string           `json:"sticker_set_name"`
	CanSetStickerSet bool             `json:"can_set_sticker_set"`
}

type ChatPermissions struct {
}

type ChatPhoto struct {
}

// Message represents a message
type Message struct {
	MessageID             int                   `json:"message_id"`
	From                  *User                 `json:"from"`
	Date                  int                   `json:"date"`
	Chat                  *Chat                 `json:"chat"`
	ForwardFrom           *User                 `json:"forward_from"`
	ForwardFromChat       *Chat                 `json:"forward_from_chat"`
	ForwardFromMessageID  int                   `json:"forward_from_message_id"`
	ForwardSignature      string                `json:"forward_signature"`
	ForwardSenderName     string                `json:"forward_sender_name"`
	ForwardDate           int                   `json:"forward_date"`
	ReplyToMessage        *Message              `json:"reply_to_message"`
	ViaBot                *User                 `json:"via_bot"`
	EditDate              int                   `json:"edit_date"`
	MediaGroupID          string                `json:"media_group_id"`
	AuthorSignature       string                `json:"author_signature"`
	Text                  string                `json:"text"`
	Entities              *[]MessageEntity      `json:"entities"`
	Animation             *Animation            `json:"animation"`
	Audio                 *Audio                `json:"audio"`
	Document              *Document             `json:"document"`
	Photo                 *[]PhotoSize          `json:"photo"`
	Sticker               *Sticker              `json:"sticker"`
	Video                 *Video                `json:"video"`
	VideoNote             *VideoNote            `json:"video_note"`
	Voice                 *Voice                `json:"voice"`
	Caption               string                `json:"caption"`
	CaptionEntities       *[]MessageEntity      `json:"caption_entities"`
	Contact               *Contact              `json:"contact"`
	Dice                  *Dice                 `json:"dice"`
	Game                  *Game                 `json:"game"`
	Poll                  *Poll                 `json:"poll"`
	Venue                 *Venue                `json:"venute"`
	Location              *Location             `json:"location"`
	NewChatMembers        *[]User               `json:"new_chat_members"`
	LeftChatMember        *User                 `json:"left_chat_member"`
	NewChatTitle          string                `json:"new_chat_title"`
	NewChatPhoto          *[]PhotoSize          `json:"new_chat_photo"`
	DeleteChatPhoto       bool                  `json:"delete_chat_photo"`
	GroupChatCreated      bool                  `json:"group_chat_created"`
	SupergroupChatCreated bool                  `json:"supergroup_chat_created"`
	ChannelChatCreated    bool                  `json:"channel_chat_created"`
	MigrateToChatID       int                   `json:"migrate_to_chat_id"`
	MigrateFromChatID     int                   `json:"migrate_from_chat_id"`
	PinnedMessage         *Message              `json:"pinned_message"`
	ConnectedWebsite      string                `json:"connected_website"`
	PassportData          *PassportData         `json:"passport_data"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
}

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	URL      string `json:"url"`
	User     *User  `json:"user"`
	Language string `json:"language"`
}

type Animation struct {
}

type Audio struct {
}

type Document struct {
}

type PhotoSize struct {
}

type Sticker struct {
}

type Video struct {
}

type VideoNote struct {
}

type Voice struct {
}

type Contact struct {
}

type Dice struct {
}

type Game struct {
}

type Venue struct {
}

type Location struct {
}

type Invoce struct {
}

type PassportData struct {
}

type InlineKeyboardMarkup struct {
}

// IsCommand returns true if message starts with a "bot_command" entity.
func (m *Message) IsCommand() bool {
	if m.Entities == nil || len(*m.Entities) == 0 {
		return false
	}

	entity := (*m.Entities)[0]
	return entity.Offset == 0 && entity.Type == "bot_command"
}

// Command checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
func (m *Message) Command() string {
	if !m.IsCommand() {
		return ""
	}

	entity := (*m.Entities)[0]
	command := m.Text[1:entity.Length]

	if i := strings.Index(command, "@"); i != -1 {
		command = command[:i]
	}

	return command
}
