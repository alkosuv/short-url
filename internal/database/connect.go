package database

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var (
	errNoSuchTableURL = errors.New("no such table: url")
)

func NewConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./sqlite/short-url.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := checkTable(db); err != nil {
		return nil, err
	}

	return db, nil
}

func checkTable(db *sql.DB) (err error) {
	defer func() {
		recover()

		if err != nil && err.Error() == errNoSuchTableURL.Error() {
			err = createTable(db)
		}
	}()

	query := "SELECT COUNT(1) FROM url"
	if _, err = db.Exec(query); err != nil {
		return
	}

	return
}

func createTable(db *sql.DB) (err error) {
	query := `
		CREATE TABLE url (
			id INTEGER PRIMARY KEY,
			original TEXT NOT NULL,
			hash TEXT NOT NULL
		);
	`

	_, err = db.Exec(query)
	return
}
