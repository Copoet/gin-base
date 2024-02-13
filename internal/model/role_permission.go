package model

import "gorm.io/gorm"

/*
*
CREATE TABLE `auht_role_permission` (

	`id` int NOT NULL AUTO_INCREMENT,
	`role_id` int NOT NULL COMMENT '角色ID',
	`permisson_id` int DEFAULT NULL COMMENT '权限ID',
	`status` tinyint(1) DEFAULT NULL COMMENT '是否启用 1是2否',
	`is_delete` tinyint(1) DEFAULT NULL COMMENT '是否删除 1 是 2 否',
	`create_time` datetime DEFAULT NULL COMMENT '创建时间',
	`update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`) USING BTREE

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='角色权限关联表';
*/
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
