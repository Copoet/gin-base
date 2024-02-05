package model

import (
	"gorm.io/gorm"
)

/*
*
CREATE TABLE `system` (

	`id` mediumint NOT NULL AUTO_INCREMENT COMMENT 'id',
	`sys_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '参数名称',
	`sys_value` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '参数值',
	`sys_explain` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '参数说明',
	`sys_type` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '参数类型',
	`create_time` datetime DEFAULT NULL COMMENT '添加时间',
	`update_time` datetime DEFAULT NULL COMMENT '更新时间',
	`status` tinyint(1) DEFAULT NULL COMMENT '状态 1 启用 2 禁用',
	`is_delete` tinyint(1) DEFAULT '2' COMMENT '是否删除 1 是 2 否',
	PRIMARY KEY (`id`) USING BTREE

) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3 COMMENT='系统参数表';
*/
type System struct {
	BaseModel
	ID         int    `json:"id"`
	SysName    string `json:"sys_name"`
	SysValue   string `json:"sys_value"`
	SysExplain string `json:"sys_explain"`
	SysType    string `json:"sys_type"`
	IsDelete   int    `json:"is_delete"`
	Status     int    `json:"status"`
}

func (*System) TableName() string {
	return "system"
}

type SystemQuery struct {
	IsDelete *int
	SysName  *string
	SysType  *string
	SysValue *string
	Status   *int
}

func buildSystemQuery(dbQuery *gorm.DB, query SystemQuery) *gorm.DB {
	if query.IsDelete != nil {
		dbQuery = dbQuery.Where("is_delete = ?", *query.IsDelete)
	}
	if query.SysName != nil {
		dbQuery = dbQuery.Where("sys_name = ?", *query.SysName)
	}
	if query.SysType != nil {
		dbQuery = dbQuery.Where("sys_type = ?", *query.SysType)
	}
	if query.SysValue != nil {
		dbQuery = dbQuery.Where("sys_value = ?", *query.SysValue)
	}
	if query.Status != nil {
		dbQuery = dbQuery.Where("status = ?", *query.Status)
	}
	return dbQuery
}

func GetSystemList(query SystemQuery, page int, page_size int) (map[string]interface{}, error) {
	dbQuery := buildSystemQuery(DB.Model(&System{}), query)
	var total int64
	result := make([]System, 0)
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

func GetSystemInfo(query SystemQuery) (system *System, err error) {
	dbQuery := buildSystemQuery(DB.Model(&System{}), query)
	err = dbQuery.First(&system).Error
	if err != nil {
		return nil, err
	}
	return system, nil
}

func AddSystem(system *System) (rid int, err error) {
	result := DB.Model(&System{}).Create(system)
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}

func UpdateSystem(id int, system *System) (rid int, err error) {
	result := DB.Model(&System{}).Where("id = ?", id).Updates(system)
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}

func DeleteSystem(id int) (rid int, err error) {
	result := DB.Model(&System{}).Where("id = ?", id).Delete(&System{})
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}
