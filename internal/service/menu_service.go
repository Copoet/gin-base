package services

import (
	"gin-base/internal/model"
	"time"
)

type MenuService struct {
}

// 根据指定条件获取列表
func (s *MenuService) GetMenuList(query model.MenuQuery, page int, pageSize int) (map[string]interface{}, error) {
	data, err := model.GetMenuList(query, page, pageSize)
	return data, err
}

// 获取单条信息
func (s *MenuService) GetMenuInfo(query model.MenuQuery) (*model.Menu, error) {
	data, err := model.GetMenuInfo(query)
	return data, err
}

// 新增
func (s *MenuService) AddMenu(data *model.Menu) (id int, err error) {
	date := time.Now().Format("2006-01-02 15:04:05")
	if data.CreateTime == "" {
		data.CreateTime = date
	}
	if data.UpdateTime == "" {
		data.UpdateTime = date
	}
	id, err = model.AddMenu(data)
	return id, err
}

// 更新
func (s *MenuService) UpdateMenu(id int, data *model.Menu) (rid int, err error) {

	rid, err = model.UpdateMenu(id, data)
	if err != nil {
		return 0, err
	}
	return rid, nil
}

// 删除
func (s *MenuService) DeleteMenu(id int) (rid int, err error) {
	rid, err = model.DeleteMenu(id)
	if err != nil {
		return 0, err
	}
	return rid, nil
}

// 获取所有菜单
func (s *MenuService) GetAllMenu() ([]*model.Menu, error) {

	data, err := model.GetAllMenu()
	return data, err
}
