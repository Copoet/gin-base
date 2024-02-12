package model

type Role struct {
	BaseModel
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Status   int    `json:"status"`
	IsDelete int    `json:"is_delete"`
}

// 指定表名
func (Role) TableName() string {
	return "auth_role"
}

type QueryRole struct {
	Name   *string
	ID     *int
	Status *int
}
