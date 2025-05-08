package models

import "api/db"

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
