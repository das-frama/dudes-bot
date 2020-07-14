package command

import "database/sql"

// RowQueryer models part of a database/sql.DB.
type RowQueryer interface {
	QueryRow(string, ...interface{}) *sql.Row
}
