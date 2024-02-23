package services

import "gin-base/internal/model"

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
