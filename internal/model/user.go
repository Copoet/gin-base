package model

// User 用户模型
type User struct {
	BaseModel
	Username string
	Sex      string
	Age      int
}
