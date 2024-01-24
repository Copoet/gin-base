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
