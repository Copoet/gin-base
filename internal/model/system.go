package model

import (
	"gorm.io/gorm"
)

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
