package services

import (
	"gin-base/internal/model"
)

type ManagerService struct {
}

// 根据指定条件获取列表
func (s *ManagerService) GetManagerList(query model.ManagerQuery, page int, pageSize int) (map[string]interface{}, error) {

	data, err := model.GetManagerList(query, page, pageSize)
	return data, err
}

// 获取单条信息
func (s *ManagerService) GetManagerInfo(query model.ManagerQuery) (*model.AuthManager, error) {
	data, err := model.GetManagerInfo(query)
	return data, err
}
