package sqlite

import (
	"database/sql"
	"io/ioutil"
)

// Migrate is read init sql file.
func Migrate(db *sql.DB, path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Exec sql.
	_, err = db.Exec(string(file))
	if err != nil {
		return err
	}

	return nil
}
