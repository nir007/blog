package main

import (
	"database/sql"
	"github.com/spf13/viper"
	"log"
)

func getDatabase() (db *sql.DB) {
	connection := viper.GetString("db.connection")

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
