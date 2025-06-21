package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB(connect_string string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connect_string)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Ping(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}
