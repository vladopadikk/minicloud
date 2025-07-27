package service

import (
	"errors"
	"fmt"
	"minicloud/model"
	"minicloud/storage"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func Register(user model.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("username and password required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = storage.CreateUser(user.Username, string(hashedPassword))
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return errors.New("user already exists")
		}
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func Login(req model.User) (string, model.User, error) {
	user, err := storage.GetUserByUsername(req.Username)

	if err != nil {
		return "", model.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", model.User{}, errors.New("invalid password")
	}

	token := uuid.NewString()
	lifeTime := time.Now().Add(24 * time.Hour)

	err = storage.SaveSession(user.ID, token, lifeTime)
	if err != nil {
		return "", model.User{}, fmt.Errorf("failed to save session: %w", err)
	}

	return token, user, nil
}
