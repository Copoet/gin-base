package services

import (
	"gin-base/internal/model"
)

type ManagerService struct {
}

func (s *ManagerService) GetManagerList(query model.ManagerQuery, page int, pageSize int) (map[string]interface{}, error) {

	data, err := model.GetManagerList(query, page, pageSize)
	return data, err
}
