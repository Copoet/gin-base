package services

import (
	"gin-base/internal/model"
)

type UserService struct {
}

func (u *UserService) GetUsers(page int, pageSize int, name string) (data map[string]interface{}, err error) {
	var query model.UserQuery

	query.Keyword = &name

	res, err := model.GetUserList(query, page, pageSize)
	if err != nil {
		return nil, err
	}
	return res, nil
}
