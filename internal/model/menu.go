package model

import "gorm.io/gorm"

type Menu struct {
	BaseModel
	ID       int    `json:"id"`
	Status   int    `json:"status"`
	Path     string `json:"path"`
	ParentId int    `json:"parent_id"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	IsDelete int    `json:"is_delete"`
}

func (*Menu) TableName() string {
	return "menu"
}

type MenuQuery struct {
	ParentID int
	ID       *int
	ParentId *int
	IsDelete *int
	Name     *string
	Status   *int
}

func buildMenuQuery(db *gorm.DB, query MenuQuery) *gorm.DB {
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
	return db
}

func GetMenuList(query MenuQuery, page int, page_size int) (map[string]interface{}, error) {
	var menus []*Menu
	var total int64
	dbQuery := buildMenuQuery(DB.Model(&Menu{}), query)
	//统计条数
	err := dbQuery.Count(&total).Error
	if err != nil {
		return nil, err
	}
	//查询数据
	err = dbQuery.Offset((page - 1) * page_size).Limit(page_size).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"total": total,
		"list":  menus,
	}, nil
}

func GetMenuInfo(query MenuQuery) (menu *Menu, err error) {
	dbQuery := buildMenuQuery(DB.Model(&Menu{}), query)
	err = dbQuery.First(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func AddMenu(menu *Menu) (rid int, err error) {
	result := DB.Model(&Menu{}).Create(menu)
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}

func UpdateMenu(id int, menu *Menu) (rid int, err error) {
	result := DB.Model(&Menu{}).Where("id = ?", id).Updates(menu)
	if result.Error != nil {
		return 0, result.Error
	}
	return rid, nil
}

func DeleteMenu(id int) (rid int, err error) {
	result := DB.Model(&Menu{}).Where("id = ?", id).Delete(&Menu{})
	if result.Error != nil {
		return 0, result.Error
	}
	return 0, nil
}

// 获取所有菜单
func GetAllMenu() ([]*Menu, error) {
	var menus []*Menu
	err := DB.Model(&Menu{}).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}
