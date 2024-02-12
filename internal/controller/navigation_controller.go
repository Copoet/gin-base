package controllers

import (
	"gin-base/internal/model"
	services "gin-base/internal/service"
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type NavigationController struct {
	BaseController
	NavigationService *services.NavigationService
}

// @Summary 获取导航列表
// @Tags Navigation
// @Produce  json
// @Param name query string false "name"
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Router /nav/list [get]
func (n *NavigationController) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	name := c.Query("name")
	status, _ := strconv.Atoi(c.Query("status"))
	var query model.NavigationQuery
	if name != "" {
		query.Name = &name
	}
	if status != 0 {
		query.Status = &status
	}
	data, err := n.NavigationService.GetNavigationList(query, page, pageSize)
	if err != nil {
		n.ReturnFail(c, util.PublicError, err)
	}
	n.ReturnSuccess(c, util.PublicSuccess, data)
}

// @Summary 添加导航
// @Tags Navigation
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   name   query    string     true        "name"
// @Param   parentId   query    string     true        "parentId"
// @Param   url   query    string     true        "url"
// @Param   status   query    int     true        "status"
// @Router /nav/add [post]
func (n *NavigationController) AddNavigation(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		n.ReturnFail(c, util.PublicParamsNull, nil)
	}
	parentId, _ := strconv.Atoi(c.PostForm("parent_id"))
	url := c.PostForm("url")
	status, _ := strconv.Atoi(c.PostForm("status"))
	data, err := n.NavigationService.AddNavigation(&model.Navigation{
		Name:     name,
		Status:   status,
		Url:      url,
		ParentId: parentId,
	})
	if err != nil {
		n.ReturnFail(c, util.PublicError, err)
	}
	n.ReturnSuccess(c, util.PublicSuccess, data)
}

// @Summary 更新导航
// @Tags Navigation
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   name   query    string     true        "name"
// @Param   parentId   query    string     true        "parentId"
// @Param   url   query    string     true        "url"
// @Param   status   query    int     true        "status"
// @Router /nav/update [post]
func (n *NavigationController) UpdateNavigation(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parent_id"))
	url := c.PostForm("url")
	status, _ := strconv.Atoi(c.PostForm("status"))
	data, err := n.NavigationService.UpdateNavigation(id, &model.Navigation{
		Name:     name,
		ParentId: parentId,
		Url:      url,
		Status:   status,
	})
	if err != nil {
		n.ReturnFail(c, util.PublicError, err)
	}
	n.ReturnSuccess(c, util.PublicSuccess, data)
}

// @Summary 删除导航
// @Tags Navigation
// @Accept  x-www-form-urlencoded
// @Param   id   formData  int  true  "ID"
// @Produce  json
// @Router /nav/delete [delete]
func (n *NavigationController) DeleteNavigation(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	if id == 0 {
		n.ReturnFail(c, util.PublicParamsNull, nil)
	}
	data, err := n.NavigationService.DeleteNavigation(id)
	if err != nil {
		n.ReturnFail(c, util.PublicError, err)
	}
	n.ReturnSuccess(c, util.PublicSuccess, data)

}
