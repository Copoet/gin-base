package services

import (
	"gin-base/internal/model"
	"time"
)

type SystemService struct {
}

// 获取系统配置列表
func (s *SystemService) GetSystemList(query model.SystemQuery, page int, pageSize int) (map[string]interface{}, error) {

	data, err := model.GetSystemList(query, page, pageSize)
	return data, err
}

// 获取单条信息
func (s *SystemService) GetSystemInfo(query model.SystemQuery) (*model.System, error) {
	data, err := model.GetSystemInfo(query)
	return data, err
}

// 新增
func (s *SystemService) AddSystem(data *model.System) (id int, err error) {
	date := time.Now().Format("2006-01-02 15:04:05")
	if data.CreateTime == "" {
		data.CreateTime = date
	}
	if data.UpdateTime == "" {
		data.UpdateTime = date
	}
	rid, err := model.AddSystem(data)
	if err != nil {
		return 0, err
	}
	return rid, nil
}

// 更新
func (s *SystemService) UpdateSystem(id int, data *model.System) (rid int, err error) {

	rid, err = model.UpdateSystem(id, data)
	if err != nil {
		return 0, err
	}
	return rid, nil
}

// 删除
func (s *SystemService) DeleteSystem(id int) (rid int, err error) {
	rid, err = model.DeleteSystem(id)
	if err != nil {
		return 0, err
	}
	return rid, nil
}
