package services

import (
	"gin-base/internal/model"
	"time"
)

type ArticleService struct {
}

// 根据指定条件获取列表
func (s *ArticleService) GetArticleList(query model.ArticleQuery, page int, pageSize int) (map[string]interface{}, error) {
	data, err := model.GetArticleList(query, page, pageSize)
	return data, err
}

// 获取单条信息
func (s *ArticleService) GetArticleInfo(query model.ArticleQuery) (*model.Article, error) {
	data, err := model.GetArticleInfo(query)
	return data, err
}

// 新增文章
func (s *ArticleService) AddArticle(data *model.Article) (id int, err error) {
	date := time.Now().Format("2006-01-02 15:04:05")
	if data.CreateTime == "" {
		data.CreateTime = date
	}
	if data.UpdateTime == "" {
		data.UpdateTime = date
	}
	id, err = model.AddArticle(data)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新文章信息
func (s *ArticleService) UpdateArticle(id int, data *model.Article) (rid int, err error) {
	id, err = model.UpdateArticle(id, data)
	if err != nil {
		return 0, err
	}
	return rid, nil
}

// 删除文章
func (s *ArticleService) DeleteArticle(id int) error {
	_, err := model.DeleteArticle(id)
	if err != nil {
		return err
	}
	return nil
}
