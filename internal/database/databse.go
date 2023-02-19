package database

import (
	"database/sql"
)

type datebase struct {
	conn *sql.DB
}

type Database interface {
	GetByOriginal(original string) (hash string, err error)
	GetByHash(hash string) (original string, err error)
	Set(original, hash string) error
}

func New(conn *sql.DB) (Database, error) {
	return &datebase{conn: conn}, nil
}

func (db *datebase) GetByHash(hash string) (original string, err error) {
	query := "SELECT original FROM url WHERE hash = $1"
	err = db.conn.QueryRow(query, hash).Scan(&original)
	return
}

func (db *datebase) GetByOriginal(original string) (hash string, err error) {
	query := "SELECT hash FROM url WHERE original = $1"
	err = db.conn.QueryRow(query, original).Scan(&hash)
	return
}

func (db *datebase) Set(original, hash string) (err error) {
	query := "INSERT INTO url (original, hash) VALUES ($1, $2)"
	_, err = db.conn.Exec(query, original, hash)
	return
}
