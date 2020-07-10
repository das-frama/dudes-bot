package sqlite

import (
	"database/sql"
)

// LoadDB create/open a sqlite db file.
func LoadDB(path string) (*sql.DB, error) {
	// Open DB.
	db, err := sql.Open("sqlite3", "db/bot.db")
	if err != nil {
		return db, err
	}

	return db, nil
}
