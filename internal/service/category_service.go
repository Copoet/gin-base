package services

import "gin-base/internal/model"

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

	err = model.AddCategory(data)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新分类
func (s *CategoryService) UpdateCategory(id int, data *model.Category) (rid int, err error) {
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
