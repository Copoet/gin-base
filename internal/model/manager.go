package model

import (
	"fmt"
	"gorm.io/gorm"
)

type AuthManager struct {
	BaseModel
	ID            int    `json:"id"`
	Status        int    `json:"status"`
	Uuid          string `json:"uuid"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	RememberToken string `json:"remember_token"`
	UpIp          string `json:"up_ip"`
	Remark        string `json:"remark"`
	IsDelete      int    `json:"is_delete"`
	LastTime      string `json:"last_at"`
}

// 指定表名
func (AuthManager) TableName() string {
	return "auth_manager"
}

type ManagerQuery struct {
	Keyword  *string
	IsDelete *int
	Name     *string
	Status   *int
}

// 构造可复用的查询条件
func buildManagerQuery(db *gorm.DB, query ManagerQuery) *gorm.DB {
	if query.Keyword != nil {
		db = db.Where("name LIKE ?", "%"+*query.Keyword+"%")
	}
	if query.IsDelete != nil {
		db = db.Where("is_delete = ?", *query.IsDelete)
	}
	if query.Name != nil {
		db = db.Where("name = ?", *query.Name)
	}
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}
	return db
}

// 根据指定条件获取列表
func GetManagerList(query ManagerQuery, page int, pageSize int) (result map[string]interface{}, err error) {
	var managers []*AuthManager
	var total int64
	dbQuery := buildManagerQuery(DB.Model(&AuthManager{}), query)

	//统计条数
	err = dbQuery.Count(&total).Error
	if err != nil {
		return nil, err
	}
	//查询数据
	err = dbQuery.Offset((page - 1) * pageSize).Limit(pageSize).Find(&managers).Error
	if err != nil {
		return nil, err
	}
	result = map[string]interface{}{
		"total": total,
		"list":  managers,
	}
	return result, nil
}

// 根据指定条件获取单条信息
func GetManagerInfo(query ManagerQuery) (manager *AuthManager, err error) {
	dbQuery := buildManagerQuery(DB.Model(&AuthManager{}), query)
	err = dbQuery.First(&manager).Error
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return manager, nil

}

// 添加管理员
func AddManager(manager *AuthManager) (id int, err error) {
	result := DB.Create(manager)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(manager.ID), nil
}

// 更新管理员
func UpdateManager(id int, manager *AuthManager) (rid int, err error) {
	result := DB.Model(&AuthManager{}).Where("id = ?", id).Updates(manager)
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}

// 删除管理员
func DeleteManager(id int) (rid int, err error) {
	result := DB.Model(&AuthManager{}).Where("id = ?", id).Update("is_delete", 1)
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil

}
