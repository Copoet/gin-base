package model

// Category 文章分类
type Category struct {
	BaseModel
	ID          int    `json:"id"`
	ParentID    int    `json:"parent_id"`
	SortName    string `json:"sort_name"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	IsDelete    int    `json:"is_delete"`
}
