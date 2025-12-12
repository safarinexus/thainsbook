package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type User struct {
	Uid            string `json:"uid"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	DateJoined     string `json:"date_joined"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserModel struct {
	DB *sql.DB
}

func (u UserModel) GetUser(username string) (User, error) {
	var user User

	row := u.DB.QueryRow("SELECT id, username, password_hash, created_at FROM users WHERE username = ?", username)
	if err := row.Scan(&user.Uid, &user.Username, &user.HashedPassword, &user.DateJoined); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("Error in fetching user: %s", err)
	}

	return user, nil
}

func (u UserModel) GetUserPassword(username string) (string, error) {
	var hashedPassword string

	row := u.DB.QueryRow("SELECT password_hash FROM users WHERE username = ?", username)
	if err := row.Scan(&hashedPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found")
		}
		return "", fmt.Errorf("Error in fetching user password: %s", err)
	}

	return hashedPassword, nil
}

func (u UserModel) AddUser(uid string, username string, hashedPassword string) error {
	res, err := u.DB.Exec("INSERT INTO users (id, username, password_hash) VALUES (?, ?, ?)", uid, username, hashedPassword)
	if err != nil {
		return fmt.Errorf("Error in adding user: %s", err)
	}
	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error in adding user: %s", err)
	}
	log.Printf("User added: %s", username)
	return nil
}
