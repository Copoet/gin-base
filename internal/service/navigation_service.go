package services

import (
	"gin-base/internal/model"
	"time"
)

type NavigationService struct{}

// 获取导航列表
func (s *NavigationService) GetNavigationList(query model.NavigationQuery, page int, pageSize int) (map[string]interface{}, error) {
	data, err := model.GetNavigationList(query, page, pageSize)
	return data, err
}

// 获取单条信息
func (s *NavigationService) GetNavigationInfo(query model.NavigationQuery) (*model.Navigation, error) {
	data, err := model.GetNavigationInfo(query)
	return data, err
}

// 新增
func (s *NavigationService) AddNavigation(data *model.Navigation) (id int, err error) {
	date := time.Now().Format("2006-01-02 15:04:05")
	if data.CreateTime == "" {
		data.CreateTime = date
	}
	if data.UpdateTime == "" {
		data.UpdateTime = date
	}
	id, err = model.AddNavigation(data)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新
func (s *NavigationService) UpdateNavigation(id int, data *model.Navigation) (rid int, err error) {
	if data.UpdateTime == "" {
		data.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	}
	rid, err = model.UpdateNavigation(id, data)
	if err != nil {
		return 0, err
	}
	return rid, nil
}

// 删除
func (s *NavigationService) DeleteNavigation(id int) (rid int, err error) {

	rid, err = model.DeleteNavigation(id)
	if err != nil {
		return 0, err
	}
	return rid, nil
}
