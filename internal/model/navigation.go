package model

import (
	"gorm.io/gorm"
)

type Navigation struct {
	BaseModel
	ID       int    `json:"id"`
	Status   int    `json:"status"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	ParentId int    `json:"parent_id"`
	IsDelete int    `json:"is_delete"`
}

// 指定表名
func (*Navigation) TableName() string {
	return "navigation"
}

type NavigationQuery struct {
	ID       *int
	ParentId *int
	IsDelete *int
	Name     *string
	Status   *int
	Url      *string
}

func buildNavigationQuery(db *gorm.DB, query NavigationQuery) *gorm.DB {
	if query.ID != nil {
		db = db.Where("id = ?", *query.ID)
	}
	if query.ParentId != nil {
		db = db.Where("parent_id = ?", *query.ParentId)
	}
	if query.IsDelete != nil {
		db = db.Where("is_delete = ?", *query.IsDelete)
	}
	if query.Name != nil {
		db = db.Where("name = ?", *query.Name)
	}
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}
	if query.Url != nil {
		db = db.Where("url = ?", *query.Url)
	}
	return db
}

func GetNavigationList(query NavigationQuery, page int, page_size int) (map[string]interface{}, error) {
	dbQuery := buildNavigationQuery(DB.Model(&Navigation{}), query)
	var total int64
	result := make([]Navigation, 0)
	err := dbQuery.Count(&total).Error
	if err != nil {
		return nil, err
	}
	err = dbQuery.Offset((page - 1) * page_size).Limit(page_size).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"total": total,
		"list":  result,
	}, nil
}

func GetNavigationInfo(query NavigationQuery) (*Navigation, error) {
	dbQuery := buildNavigationQuery(DB.Model(&Navigation{}), query)
	var result Navigation
	err := dbQuery.First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func AddNavigation(navigation *Navigation) (rid int, err error) {
	result := DB.Model(&Navigation{}).Create(navigation)
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}

func DeleteNavigation(id int) (rid int, err error) {
	result := DB.Model(&Navigation{}).Where("id = ?", id).Delete(&Navigation{})
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}

func UpdateNavigation(id int, navigation *Navigation) (rid int, err error) {
	result := DB.Model(&Navigation{}).Where("id = ?", id).Updates(navigation)
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}
