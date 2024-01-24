package model

// User 用户模型
type Users struct {
	BaseModel
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	IsVip    int    `json:"is_vip"`
}

func GetUserList(page int, pageSize int, name string) ([]*Users, int64, error) {
	var users []*Users
	var count int64

	offset := (page - 1) * pageSize
	query := DB.Model(&Users{})

	if name != "" {
		query = query.Where("username LIKE ?", "%"+name+"%")
	}
	err := query.Count(&count).Limit(pageSize).Offset(offset).Find(&users).Error
	return users, count, err
}
