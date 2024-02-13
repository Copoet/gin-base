package model

import "gorm.io/gorm"

type RolePermission struct {
	BaseModel
	ID           int `json:"id"`
	RoleID       int `json:"role_id"`
	PermissionID int `json:"permission_id"`
	IsDelete     int `json:"is_delete"`
	Status       int `json:"status"`
}

type QueryRolePermission struct {
	RoleID       *int
	PermissionID *int
	IsDelete     *int
	Status       *int
}

func buildRolePermissionQuery(db *gorm.DB, query *QueryRolePermission) *gorm.DB {
	if query.RoleID != nil {
		db = db.Where("role_id = ?", *query.RoleID)
	}
	if query.PermissionID != nil {
		db = db.Where("permission_id = ?", *query.PermissionID)
	}
	if query.IsDelete != nil {
		db = db.Where("is_delete = ?", *query.IsDelete)
	}
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}
	return db
}

func GetRolePermissionList(query *QueryRolePermission) ([]*RolePermission, error) {
	db := DB.Model(&RolePermission{})
	db = buildRolePermissionQuery(db, query)
	var list []*RolePermission
	err := db.Find(&list).Error
	return list, err
}

func AddRolePermission(rolePermission *RolePermission) (id int, err error) {
	err = DB.Create(rolePermission).Error
	return rolePermission.ID, err
}

func DeleteRolePermission(id int) error {
	return DB.Where("id = ?", id).Delete(&RolePermission{}).Error
}

func UpdateRolePermission(rolePermission *RolePermission) error {
	return DB.Model(&RolePermission{}).Where("id = ?", rolePermission.ID).Updates(rolePermission).Error
}
