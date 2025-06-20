package repository

import "database/sql"

func Ping(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		panic("cannot read config: " + err.Error())
	}
	return nil
}
