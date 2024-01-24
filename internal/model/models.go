package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
}

// Create 在数据库中创建记录
func (base *BaseModel) Create(db *gorm.DB) error {
	return db.Create(base).Error
}

// Update 更新数据库中的记录
func (base *BaseModel) Update(db *gorm.DB, values interface{}) error {
	return db.Model(base).Updates(values).Error
}

// Delete 从数据库中删除记录
func (base *BaseModel) Delete(db *gorm.DB) error {
	return db.Delete(base).Error
}
