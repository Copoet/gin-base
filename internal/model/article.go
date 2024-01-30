package model

type Article struct {
	BaseModel
	ID          int    `json:"id"`
	Status      int    `json:"status"`
	SortID      int    `json:"sort_id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	TitlePic    string `json:"title_pic"`
	Flag        string `json:"flag"`
	Content     string `json:"content"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Views       int    `json:"views"`
	IsDelete    int    `json:"is_delete"`
}
