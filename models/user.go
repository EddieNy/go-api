package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO users(email,password) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	user.ID = userId
	return err
}

func (user User) Validate() error {
	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	validPassword := utils.CheckPassword(user.Password, retrievedPassword)

	if !validPassword {
		return errors.New("Credentials invalid")
	}

	return nil
}
