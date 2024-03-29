package controllers

import (
	"gin-base/internal/model"
	services "gin-base/internal/service"
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CategoryController struct {
	BaseController
	CategoryService *services.CategoryService
}

// @Summary 分类列表
// @Tags Category
// @Produce  json
// @Param keyword query string false "keyword"
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Router /category/list [get]
func (c *CategoryController) GetList(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	keyword := ctx.Query("keyword")
	data, err := c.CategoryService.GetCategoryList(model.CategoryQuery{
		Keyword: &keyword,
	}, page, pageSize)
	if err != nil {
		c.ReturnFail(ctx, util.PublicError, nil)
	}
	c.ReturnSuccess(ctx, util.PublicSuccess, data)
}

// @Summary 分类详情
// @Tags Category
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   id   path    int     true        "Category ID"
// @Router /category/info [get]
func (c *CategoryController) GetCategoryInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var query model.CategoryQuery
	query.Id = &id
	data, err := c.CategoryService.GetCategoryInfo(query)
	if err != nil {
		c.ReturnFail(ctx, util.PublicError, nil)
	}
	c.ReturnSuccess(ctx, util.PublicSuccess, data)
}

// @Summary 添加分类
// @Tags Category
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   name  formData  string  true  "分类名称"
// @Param   status formData  int  true  "状态" value 1 启用 0 禁用
// @Router /category/add [post]
func (c *CategoryController) AddCategory(ctx *gin.Context) {

	name := ctx.PostForm("name")
	status, _ := strconv.Atoi(ctx.PostForm("status"))
	parentId, _ := strconv.Atoi(ctx.PostForm("parent_id"))
	keywords := ctx.PostForm("keywords")
	description := ctx.PostForm("description")
	data, err := c.CategoryService.AddCategory(&model.Category{
		SortName:    name,
		Status:      status,
		ParentID:    parentId,
		Keywords:    keywords,
		Description: description,
	})
	if err != nil {
		c.ReturnFail(ctx, util.PublicError, nil)
	}
	c.ReturnSuccess(ctx, util.PublicSuccess, data)
}

// @Summary 更新分类
// @Tags Category
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   id   formData  int  true  "分类ID"
// @Param   name  formData  string  true  "分类名称"
// @Param   status formData  int  true  "状态" value 1 启用 0 禁用
// @Router /category/update [put]
func (c *CategoryController) UpdateCategory(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	status, _ := strconv.Atoi(ctx.PostForm("status"))
	parentId, _ := strconv.Atoi(ctx.PostForm("parent_id"))
	keywords := ctx.PostForm("keywords")
	description := ctx.PostForm("description")
	data, err := c.CategoryService.UpdateCategory(id, &model.Category{
		SortName:    name,
		Status:      status,
		ParentID:    parentId,
		Keywords:    keywords,
		Description: description,
	})
	if err != nil {
		c.ReturnFail(ctx, util.PublicError, nil)
	}
	c.ReturnSuccess(ctx, util.PublicSuccess, data)
}

// @Summary 删除分类
// @Tags Category
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   id   formData  int  true  "分类ID"
// @Router /category/delete [delete]
func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	err := c.CategoryService.DeleteCategory(id)
	if err != nil {
		c.ReturnFail(ctx, util.PublicError, nil)
		return
	}
	c.ReturnSuccess(ctx, util.PublicSuccess, nil)
}

// @Summary 获取栏目树形分类
// @Tags Category
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Router /category/tree [get]
func (c *CategoryController) GetTree(ctx *gin.Context) {
	data, err := c.CategoryService.GetAllCategory()
	if err != nil {
		c.ReturnFail(ctx, util.PublicError, nil)
		return
	}
	var newNode []*util.Node
	for _, v := range data {
		n := &util.Node{
			Id:       v.ID,
			Name:     v.SortName,
			ParentId: v.ParentID,
			Children: make([]*util.Node, 0),
		}
		newNode = append(newNode, n)
	}
	//调用util中的GetTree方法
	res := util.GetTree(newNode, 0)
	c.ReturnSuccess(ctx, util.PublicSuccess, res)

}
