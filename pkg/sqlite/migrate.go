package sqlite

import (
	"database/sql"
	"io/ioutil"
	"log"
)

// Migrate
func Migrate(db *sql.DB, path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Exec sql.
	result, err := db.Exec(string(file))
	if err != nil {
		return err
	}
	log.Println(result)

	return nil
}
