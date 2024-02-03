package model

import (
	"gorm.io/gorm"
	"time"
)

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

// 指定表名
func (Category) TableName() string {
	return "article_sort"
}

// 查询参数
type CategoryQuery struct {
	Id          *int
	Keyword     *string
	SortName    *string
	ParentID    *int
	Status      *int
	Description *string
}

// 构造查询参数
func buildCategoryQuery(db *gorm.DB, query CategoryQuery) *gorm.DB {
	if query.Keyword != nil {
		db = db.Where("keyword LIKE ?", "%"+*query.Keyword+"%")
	}
	if query.SortName != nil {
		db = db.Where("sort_name = ?", *query.SortName)
	}
	if query.ParentID != nil {
		db = db.Where("parent_id = ?", *query.ParentID)
	}
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}
	if query.Description != nil {
		db = db.Where("description = ?", *query.Description)
	}
	return db
}

// 根据指定条件获取列表
func GetCategoryList(query CategoryQuery, page int, pageSize int) (map[string]interface{}, error) {
	var categories []*Category
	var total int64
	dbQuery := buildCategoryQuery(DB.Model(&Category{}), query)
	//统计条数
	err := dbQuery.Count(&total).Error
	if err != nil {
		return nil, err
	}
	//查询数据
	err = dbQuery.Offset((page - 1) * pageSize).Limit(pageSize).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"total": total,
		"list":  categories,
	}, nil
}

// 根据指定条件获取单条信息
func GetCategoryOne(query CategoryQuery) (*Category, error) {
	dbQuery := buildCategoryQuery(DB.Model(&Category{}), query)
	var category Category
	err := dbQuery.First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// 新增
func AddCategory(category *Category) error {
	return DB.Create(category).Error
}

// 更新
func UpdateCategory(id int, category *Category) error {
	result := DB.Model(&Category{}).Where("id = ?", id).Updates(category)
	return result.Error
}

// 删除
func DeleteCategory(id int) error {
	date := time.Now().Format("2006-01-02 15:04:05")
	result := DB.Model(&Category{}).Where("id = ?", id).Updates(map[string]interface{}{"is_delete": 1, "update_time": date})
	return result.Error
}

// 获取所有栏目
func GetAllCategory(query CategoryQuery) ([]*Category, error) {
	dbQuery := buildCategoryQuery(DB.Model(&Category{}), query)
	var categories []*Category
	err := dbQuery.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
