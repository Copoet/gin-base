package model

import (
	"gorm.io/gorm"
	"time"
)

type Manager struct {
	BaseModel
	ID            int       `json:"id"`
	Status        int       `json:"status"`
	Uuid          string    `json:"uuid"`
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	RememberToken string    `json:"remember_token"`
	UpIp          string    `json:"up_ip"`
	Remark        string    `json:"remark"`
	IsDelete      int       `json:"is_delete"`
	CreateAt      time.Time `json:"create_at"`
	LastAt        time.Time `json:"last_at"`
	UpdateAt      time.Time `json:"update_at"`
}

type ManagerQuery struct {
	Keyword  *string
	IsDelete *int
	Name     *string
	Status   *int
}

// 构造可复用的查询条件
func buildQuery(db *gorm.DB, query ManagerQuery) *gorm.DB {
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
	var managers []*Manager
	var total int64
	dbQuery := buildQuery(DB.Model(&Manager{}), query)

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
