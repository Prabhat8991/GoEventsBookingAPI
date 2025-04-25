package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Underscore for imports which are not directly used but we need as part of package
)

var DB *sql.DB

func InitDB() {
	var error error

	DB, error = sql.Open("sqlite3", "api.db")

	if error != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventsTable := `
	  CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		time DATETIME NOT NULL,
		userID INTEGER
	  )
	`
	_, error := DB.Exec(createEventsTable)

	if error != nil {
		panic("Could not create events table")
	}
}
