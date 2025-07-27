package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"minicloud/db"
	"minicloud/model"
	"time"
)

func CreateUser(username, hashedPassword string) error {
	_, err := db.DB.Exec(
		`INSERT INTO users (username, password) VALUES ($1, $2)
	`, username, hashedPassword)
	return err
}

func GetUserByUsername(username string) (model.User, error) {
	var u model.User

	err := db.DB.QueryRow(
		`SELECT id, username, password FROM users WHERE username = $1 
	`, username).Scan(&u.ID, &u.Username, &u.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return u, errors.New("user not found")
		}
		return u, err
	}
	return u, nil
}

func GetUsernameByToken(token string) (string, error) {
	var username string

	err := db.DB.QueryRow(`
		SELECT username 
		FROM users u
		JOIN sessions s on u.id = s.user_id
		WHERE s.token = $1
	`, token).Scan(&username)

	if err != nil {
		return "", err
	}
	return username, nil
}

func SaveSession(userID int, token string, expiresAt time.Time) error {
	_, err := db.DB.Exec(
		`DELETE FROM sessions WHERE user_id = $1
	`, userID)

	if err != nil {
		return fmt.Errorf("failed to delete old sessions: %w", err)
	}

	_, err = db.DB.Exec(
		`INSERT INTO sessions (token, user_id, created_at, expires_at) VALUES ($1, $2, $3, $4)`,
		token,
		userID,
		time.Now(),
		expiresAt,
	)
	return err
}
