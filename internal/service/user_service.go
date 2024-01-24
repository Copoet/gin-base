package services

import (
	"gin-base/internal/model"
)

type UserService struct {
}

type UserListData struct {
	List  []*model.Users `json:"list"`
	Total int64          `json:"total"`
}

func (u *UserService) GetUsers(page int, pageSize int, name string) (*UserListData, error) {
	users, count, err := model.GetUserList(page, pageSize, name)
	if err != nil {
		return nil, err
	}
	return &UserListData{
		List:  users,
		Total: count,
	}, nil
}
