package model

import "gorm.io/gorm"

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

// 指定表名
func (Article) TableName() string {
	return "article"
}

// 查询条件
type ArticleQuery struct {
	Id          *int
	Keyword     *string
	Tittle      *string
	Description *string
}

// 构造查询条件
func buildArticleQuery(db *gorm.DB, query ArticleQuery) *gorm.DB {
	if query.Keyword != nil {
		db = db.Where("keyword LIKE ?", "%"+*query.Keyword+"%")
	}
	if query.Tittle != nil {
		db = db.Where("title = ?", *query.Tittle)
	}
	if query.Description != nil {
		db = db.Where("description = ?", *query.Description)
	}
	return db

}

// 根据指定条件获取列表
func GetArticleList(query ArticleQuery, page int, pageSize int) (map[string]interface{}, error) {
	var articles []*Article
	var total int64
	dbQuery := buildArticleQuery(DB.Model(&Article{}), query)
	//统计条数
	err := dbQuery.Count(&total).Error
	if err != nil {
		return nil, err
	}
	//查询数据
	err = dbQuery.Offset((page - 1) * pageSize).Limit(pageSize).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"total": total,
		"list":  articles,
	}, nil
}

// 根据指定条件获取单条信息
func GetArticleInfo(query ArticleQuery) (*Article, error) {
	dbQuery := buildArticleQuery(DB.Model(&Article{}), query)
	var article Article
	err := dbQuery.First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// 新增
func AddArticle(article *Article) (id int, err error) {
	result := DB.Create(article)
	if result.Error != nil {
		return 0, result.Error
	}
	return article.ID, nil
}

// 更新
func UpdateArticle(id int, article *Article) (rid int, err error) {
	result := DB.Model(&Article{}).Where("id = ?", id).Updates(article)
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}

// 删除
func DeleteArticle(id int) (rid int, err error) {
	result := DB.Model(&Article{}).Where("id = ?", id).Update("is_delete", 1)
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}
