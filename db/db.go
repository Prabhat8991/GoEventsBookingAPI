package db

import (
	"database/sql"
	"log"

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
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  email TEXT NOT NULL UNIQUE,
	  password TEXT NOT NULL
	)
	`
	_, error := DB.Exec(createUsersTable)

	if error != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
	  CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		time DATETIME NOT NULL,
		userID INTEGER,
		FOREIGN KEY(userID) REFERENCES users(id)
	  )
	`
	_, error = DB.Exec(createEventsTable)

	if error != nil {
		log.Println("DB error", error)
		panic("Could not create events table")
	}
}
