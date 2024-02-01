package controllers

import (
	"gin-base/internal/model"
	services "gin-base/internal/service"
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ArticleController struct {
	BaseController
	ArticleService *services.ArticleService
}

// @Summary 文章列表
// @Tags Article
// @Produce  json
// @Param keyword query string false "keyword"
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Router /articles/list [get]
func (a *ArticleController) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	data, err := a.ArticleService.GetArticleList(model.ArticleQuery{
		Keyword: &keyword,
	}, page, pageSize)
	if err != nil {
		a.ReturnFail(c, util.PublicError, nil)
	}
	a.ReturnSuccess(c, util.PublicSuccess, data)
}

// @Summary 添加文章
// @Tags Article
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   title  formData  string  true  "标题"
// @Param   content  formData  string  true  "内容"
// @Param   status formData  int  true  "状态" value 1 启用 0 禁用
// @Router /articles/add [post]
func (a *ArticleController) AddArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	status, _ := strconv.Atoi(c.PostForm("status"))
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sortId, _ := strconv.Atoi(c.PostForm("sort_id"))
	titlePic := c.PostForm("title_pic")
	flag := c.PostForm("flag")
	if title == "" || content == "" || sortId == 0 {
		a.ReturnFail(c, util.PublicParamsIllegal, nil)
		return
	}

	data, err := a.ArticleService.AddArticle(&model.Article{
		Title:       title,
		Content:     content,
		Status:      status,
		Keywords:    keywords,
		Description: description,
		SortID:      sortId,
		TitlePic:    titlePic,
		Flag:        flag,
	})
	if err != nil {
		a.ReturnFail(c, util.PublicError, nil)
	}
	a.ReturnSuccess(c, util.PublicSuccess, data)
}

// @Summary 更新文章
// @Tags Article
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   id   path    int     true        "Article ID"
// @Param   title  formData  string  true  "标题"
// @Param   content  formData  string  true  "内容"
// @Param   status formData  int  true  "状态" value 1 启用 0 禁用
// @Router /articles/update [put]
func (a *ArticleController) UpdateArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.PostForm("title")
	content := c.PostForm("content")
	status, _ := strconv.Atoi(c.PostForm("status"))
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sortId, _ := strconv.Atoi(c.PostForm("sort_id"))
	titlePic := c.PostForm("title_pic")
	flag := c.PostForm("flag")
	isDelete, _ := strconv.Atoi(c.PostForm("is_delete"))
	if title == "" || content == "" || sortId == 0 {
		a.ReturnFail(c, util.PublicParamsIllegal, nil)
		return
	}
	data, err := a.ArticleService.UpdateArticle(id, &model.Article{
		Title:       title,
		Content:     content,
		Status:      status,
		Keywords:    keywords,
		Description: description,
		SortID:      sortId,
		TitlePic:    titlePic,
		Flag:        flag,
		IsDelete:    isDelete,
	})
	if err != nil {
		a.ReturnFail(c, util.PublicError, nil)
	}
	a.ReturnSuccess(c, util.PublicSuccess, data)
}

// @Summary 删除文章
// @Tags Article
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   id   path    int     true        "Article ID"
// @Router /articles/delete [delete]
func (a *ArticleController) DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := a.ArticleService.DeleteArticle(id)
	if err != nil {
		a.ReturnFail(c, util.PublicError, nil)
		return
	}
	a.ReturnSuccess(c, util.PublicSuccess, nil)
}

// @Summary 文章详情
// @Tags Article
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   id   path    int     true        "Article ID"
// @Router /articles/info [get]
func (a *ArticleController) GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var query model.ArticleQuery
	query.Id = &id
	data, err := a.ArticleService.GetArticleInfo(query)
	if err != nil {
		a.ReturnFail(c, util.PublicError, nil)
	}
	a.ReturnSuccess(c, util.PublicSuccess, data)
}
