package models

import (
	"api/db"
	"errors"
	"log"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) CreateUser() error {
	query := `
	  INSERT INTO users (email, password)
	  VALUES (?, ?)
	`
	result, err := db.DB.Exec(query, user.Email, user.Password)

	if err != nil {
		return err
	}

	id, error := result.LastInsertId()

	user.ID = id

	return error
}

func (user User) ValidateCredentials() error {
	query := `
     SELECT password from users WHERE email = ?
   `
	row := db.DB.QueryRow(query, user.Email)

	var retrievePassword string

	err := row.Scan(&retrievePassword)

	log.Println("Error scan", err)

	if err != nil {
		return errors.New("credentials invalid")
	}

	if user.Password != retrievePassword {
		log.Println("user.Password != retrievePassword", user.Password != retrievePassword)
		return errors.New("credentials invalid")
	}

	return nil
}
