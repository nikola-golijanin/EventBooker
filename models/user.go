package models

import (
	"errors"
	"homelab/event-booker/db"
	"homelab/event-booker/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users (email, password) VALUES (?, ?)`

	sql, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer sql.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := sql.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := `
	SELECT password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil
}
