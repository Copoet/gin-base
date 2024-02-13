package model

import "gorm.io/gorm"

type ManagerRole struct {
	BaseModel
	ID        int `json:"id"`
	ManagerID int `json:"manager_id"`
	RoleID    int `json:"role_id"`
	IsDelete  int `json:"is_delete"`
	Status    int `json:"status"`
}

func (ManagerRole) TableName() string {
	return "auth_manager_role"
}

type QueryManagerRole struct {
	ID        *int
	ManagerID *int
	RoleID    *int
	IsDelete  *int
	Status    *int
}

func buildManagerRoleQuery(db *gorm.DB, query *QueryManagerRole) *gorm.DB {
	if query.ID != nil {
		db = db.Where("id = ?", *query.ID)
	}
	if query.ManagerID != nil {
		db = db.Where("manager_id = ?", *query.ManagerID)
	}
	if query.RoleID != nil {
		db = db.Where("role_id = ?", *query.RoleID)
	}
	if query.IsDelete != nil {
		db = db.Where("is_delete = ?", *query.IsDelete)
	}
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}
	return db
}
