package services

import (
	"gin-base/internal/model"
	"time"
)

type CategoryService struct {
}

// 根据指定条件获取列表
func (s *CategoryService) GetCategoryList(query model.CategoryQuery, page int, pageSize int) (map[string]interface{}, error) {

	data, err := model.GetCategoryList(query, page, pageSize)
	return data, err
}

// 获取单条信息
func (s *CategoryService) GetCategoryInfo(query model.CategoryQuery) (*model.Category, error) {

	data, err := model.GetCategoryOne(query)
	return data, err
}

// 新增分类
func (s *CategoryService) AddCategory(data *model.Category) (id int, err error) {
	date := time.Now().Format("2006-01-02 15:04:05")
	if data.CreateTime == "" {
		data.CreateTime = date
	}
	if data.UpdateTime == "" {
		data.UpdateTime = date
	}
	err = model.AddCategory(data)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新分类
func (s *CategoryService) UpdateCategory(id int, data *model.Category) (rid int, err error) {
	date := time.Now().Format("2006-01-02 15:04:05")
	if data.UpdateTime == "" {
		data.UpdateTime = date
	}
	err = model.UpdateCategory(id, data)
	if err != nil {
		return 0, err
	}
	return rid, nil
}

// 删除
func (s *CategoryService) DeleteCategory(id int) error {

	err := model.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}

// 获取所有分类
func (s *CategoryService) GetAllCategory() ([]*model.Category, error) {
	data, err := model.GetAllCategory(model.CategoryQuery{})
	return data, err
}
