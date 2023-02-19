package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
}

func (user *User) Prepare(isARegister bool) error {
	if err := user.validate(isARegister); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(isARegister bool) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Nickname == "" {
		return errors.New("nickname is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if isARegister && user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Name)
}
