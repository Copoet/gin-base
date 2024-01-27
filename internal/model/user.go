package model

import (
	"gorm.io/gorm"
	"time"
)

// User 用户模型
type Users struct {
	BaseModel
	ID       int       `json:"id"`
	Phone    int       `json:"phone"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	IsVip    int       `json:"is_vip"`
	Password string    `json:"password"`
	Token    string    `json:"token"`
	Status   int       `json:"status"`
	Uuid     string    `json:"uuid"`
	RegIp    string    `json:"reg_ip"`
	IsDelete int       `json:"is_delete"`
	LastTime time.Time `json:"last_time"`
	RegTime  time.Time `json:"reg_time"`
}

// UserQuery 用户查询条件
type UserQuery struct {
	Keyword  *string
	IsDelete *int
	Phone    *string
	Status   *int
	Email    *string
}

// 指定表名
func (Users) TableName() string {
	return "users"
}

// 构造可复用的查询条件
func buildUserQuery(db *gorm.DB, query UserQuery) *gorm.DB {
	if query.Keyword != nil {
		db = db.Where("username LIKE ?", "%"+*query.Keyword+"%")
	}
	if query.IsDelete != nil {
		db = db.Where("is_delete = ?", *query.IsDelete)
	}
	if query.Phone != nil {
		db = db.Where("Phone = ?", *query.Phone)
	}
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}
	if query.Email != nil {
		db = db.Where("email = ?", *query.Email)
	}
	return db
}

// 根据指定条件获取列表
func GetUserList(query UserQuery, page int, pageSize int) (result map[string]interface{}, err error) {
	var users []*Users
	var total int64

	//初始化查询
	dbQuery := buildUserQuery(DB.Model(&Users{}), query)
	//统计条数
	err = dbQuery.Count(&total).Error
	if err != nil {
		return nil, err
	}
	//分页
	offset := (page - 1) * pageSize
	err = dbQuery.Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, err
	}
	result = map[string]interface{}{
		"total": total,
		"list":  users,
	}
	return result, nil
}
