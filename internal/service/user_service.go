package services

import (
	"gin-base/internal/database"
	"gin-base/internal/model"
)

type UserService struct {
}

func (u *UserService) GetUsers(page int, pageSize int, name string) ([]model.Users, int64, error) {
	var users []model.Users
	var count int64

	offset := (page - 1) * pageSize
	query := database.DB.Model(&model.Users{})
	if name != "" {
		query = query.Where("username LIKE ?", "%"+name+"%")
	}

	err := query.Count(&count).Limit(pageSize).Offset(offset).Find(&users).Error
	return users, count, err
}
