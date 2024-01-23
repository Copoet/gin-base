package model

// User 用户模型
type Users struct {
	BaseModel
	Username string
	Sex      string
	Age      int
}
