package services

import (
	"gin-base/internal/model"
	"time"
)

type RoleService struct {
}

// 角色列表
func (r *RoleService) GetRoleList(query model.QueryRole, page int, pageSize int) (map[string]interface{}, error) {
	data, err := model.GetRoleList(query, page, pageSize)
	return data, err
}

// 获取单条信息
func (r *RoleService) GetRoleInfo(query model.QueryRole) (*model.Role, error) {
	data, err := model.GetRoleOne(query)
	return data, err
}

// 新增
func (r *RoleService) AddRole(data *model.Role) (id int, err error) {
	date := time.Now().Format("2006-01-02 15:04:05")
	if data.CreateTime == "" {
		data.CreateTime = date
	}
	if data.UpdateTime == "" {
		data.UpdateTime = date
	}
	id, err = model.AddRole(data)
	return id, err
}

// 更新
func (r *RoleService) UpdateRole(id int, data *model.Role) (err error) {
	date := time.Now().Format("2006-01-02 15:04:05")
	if data.UpdateTime == "" {
		data.UpdateTime = date
	}
	err = model.UpdateRole(id, data)
	return err
}

// 删除
func (r *RoleService) DeleteRole(id int) (err error) {
	err = model.DeleteRole(id)
	return err
}
