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

func GetManagerRoleList(query *QueryManagerRole) ([]*ManagerRole, error) {
	db := DB.Model(&ManagerRole{})
	db = buildManagerRoleQuery(db, query)
	var list []*ManagerRole
	err := db.Find(&list).Error
	return list, err
}

func AddManagerRole(managerRole *ManagerRole) (id int, err error) {
	err = DB.Create(managerRole).Error
	return managerRole.ID, err
}

func UpdateMangerRole(managerRole *ManagerRole) (id int, err error) {
	err = DB.Updates(managerRole).Error
	return managerRole.ID, err
}

func DeleteManagerRole(id int) (rid int, err error) {
	result := DB.Model(&ManagerRole{}).Where("id = ?", id).Delete(&ManagerRole{})
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}
