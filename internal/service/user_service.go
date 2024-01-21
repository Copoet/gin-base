package services

import (
	"errors"
	"gin-base/internal/model"
)

func CheckUser(user *model.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Age <= 0 {
		return errors.New("age must be positive")
	}
	return nil
}
