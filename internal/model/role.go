package model

import "gorm.io/gorm"

type Role struct {
	BaseModel
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Status   int    `json:"status"`
	IsDelete int    `json:"is_delete"`
}

// 指定表名
func (Role) TableName() string {
	return "auth_role"
}

type QueryRole struct {
	Name   *string
	ID     *int
	Status *int
}

func buildRoleQuery(db *gorm.DB, query QueryRole) *gorm.DB {
	if query.ID != nil {
		db = db.Where("id = ?", *query.ID)
	}
	if query.Name != nil {
		db = db.Where("name = ?", *query.Name)
	}
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}
	return db
}

func GetRoleList(query QueryRole, page int, page_size int) (map[string]interface{}, error) {
	var result []*Role
	var total int64
	dbQuery := buildRoleQuery(DB.Model(&Role{}), query)
	err := dbQuery.Count(&total).Error
	if err != nil {
		return nil, err
	}
	err = dbQuery.Offset((page - 1) * page_size).Limit(page_size).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"total": total,
		"list":  result,
	}, nil
}

func GetRoleOne(query QueryRole) (*Role, error) {
	var role Role
	err := DB.Where(buildRoleQuery(DB, query)).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func AddRole(role *Role) (id int, err error) {
	result := DB.Create(role)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(role.ID), nil
}

func UpdateRole(id int, role *Role) error {
	result := DB.Model(&Role{}).Where("id = ?", id).Updates(role)
	return result.Error
}

func DeleteRole(id int) error {
	result := DB.Model(&Role{}).Where("id = ?", id).Delete(&Role{})
	return result.Error
}
