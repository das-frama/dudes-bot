package sqlite

import (
	"das-frama/dudes-bot/pkg/bot"
	"database/sql"
	"io/ioutil"
	"strings"
)

type Queryer interface {
	InitSchema(string) error
	IsChatActive(int) (bool, error)
	GetOrCreateChat(*bot.Chat) (Chat, bool, error)
	StopChat(int) error
	QueryRandomCatJoke() (CatJoke, error)
	QueryRandomPoetry(int) (Poetry, error)
}

type db struct {
	conn *sql.DB
}

// New create/open a sqlite db file.
func New(conn *sql.DB) Queryer {
	return &db{
		conn: conn,
	}
}

// InitSchema reads init sql file.
func (db *db) InitSchema(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Exec sql.
	_, err = db.conn.Exec(string(file))
	if err != nil {
		return err
	}

	return nil
}

// IsChatActive cheks if chat is active.
func (db *db) IsChatActive(id int) (bool, error) {
	row := db.conn.QueryRow("SELECT id FROM chats WHERE id = ? AND is_active = 1", id)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

// GetOrCreateChat finds and returns a chat struct by given id.
func (db *db) GetOrCreateChat(bc *bot.Chat) (Chat, bool, error) {
	chat := Chat{
		ID:            bc.ID,
		Type:          bc.Type,
		Title:         bc.Title,
		Username:      bc.Username,
		FirstName:     bc.FirstName,
		LastName:      bc.LastName,
		Description:   bc.Description,
		PinnedMessage: "",
		IsActive:      true,
	}
	created := false

	// Query row.
	row := db.conn.QueryRow("SELECT id FROM chats WHERE id = ?", chat.ID)
	switch err := row.Scan(&chat.ID); err {
	case sql.ErrNoRows:
		// Insert new record.
		_, err := db.conn.Exec(
			"INSERT INTO chats (id, type, title, username, first_name, last_name, description, pinned_message, is_active) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
			chat.ID,
			chat.Type,
			chat.Title,
			chat.Username,
			chat.FirstName,
			chat.LastName,
			chat.Description,
			chat.PinnedMessage,
			chat.IsActive,
		)
		if err != nil {
			return chat, created, err
		}
		created = true

	case nil:
		// Update chat record.
		_, err := db.conn.Exec(
			"UPDATE chats SET type = ?, title = ?, username = ?, first_name = ?, last_name = ?, description = ?, pinned_message = ? WHERE id = ?",
			chat.Type,
			chat.Title,
			chat.Username,
			chat.FirstName,
			chat.LastName,
			chat.Description,
			chat.PinnedMessage,
			chat.ID,
		)
		if err != nil {
			return chat, created, err
		}
		created = false

	default:
		return chat, created, err
	}

	return chat, created, nil
}

func (db *db) StopChat(id int) error {
	_, err := db.conn.Exec("UPDATE chats SET is_active = 0")
	return err
}

func (db *db) QueryRandomCatJoke() (CatJoke, error) {
	var joke CatJoke

	// Prepare row.
	row := db.conn.QueryRow("SELECT * FROM cat_jokes ORDER BY RANDOM() LIMIT 1")
	if err := row.Scan(&joke.ID, &joke.Text, &joke.Day); err != nil {
		return joke, err
	}

	return joke, nil
}

func (db *db) QueryRandomPoetry(t int) (Poetry, error) {
	var poetry Poetry

	// Prepare row.
	row := db.conn.QueryRow("SELECT * FROM poetry WHERE type=? ORDER BY RANDOM() LIMIT 1", t)
	if err := row.Scan(&poetry.ID, &poetry.Text, &poetry.Type); err != nil {
		return poetry, err
	}
	poetry.Text = strings.TrimSpace(poetry.Text)

	return poetry, nil
}
