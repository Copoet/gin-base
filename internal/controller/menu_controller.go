package controllers

import (
	"gin-base/internal/model"
	services "gin-base/internal/service"
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MenuController struct {
	BaseController
	MenuService *services.MenuService
}

// @Summary 菜单列表
// @Tags Menu
// @Produce  json
// @Param name query string false "name"
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Router /menu/list [get]
func (s *MenuController) GetList(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.Query("name")
	//id, _ := strconv.Atoi(c.Query("id"))
	data, err := s.MenuService.GetMenuList(
		model.MenuQuery{
			Name: &name,
		},
		page,
		pageSize)
	if err != nil {
		s.ReturnFail(c, util.PublicError, nil)
	}
	s.ReturnSuccess(c, util.PublicSuccess, data)
}

// @Summary 添加菜单
// @Tags Menu
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   name   query    string     true        "name"
// @Param   title   query    string     true        "title"
// @Param   pid   query    int     true        "pid"
// @Param   icon   query    string     true        "icon"
// @Param   path   query    string     true        "path"
// @Param   status   query    int     true        "status"
// @Router /menu/add [post]
func (s *MenuController) AddMenu(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		s.ReturnFail(c, util.PublicParamsNull, nil)
	}
	parentId, _ := strconv.Atoi(c.PostForm("parent_id"))

	data, err := s.MenuService.AddMenu(&model.Menu{
		Name:     name,
		ParentId: parentId,
	})
	if err != nil {
		s.ReturnFail(c, util.PublicError, nil)
	}
	s.ReturnSuccess(c, util.PublicSuccess, data)
}

// @Summary 删除菜单
// @Tags Menu
// @Accept  x-www-form-urlencoded
// @Router /menu/delete [delete]
func (s *MenuController) DeleteMenu(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	rid, err := s.MenuService.DeleteMenu(id)
	if err != nil {
		s.ReturnFail(c, util.PublicError, nil)
		return
	}
	s.ReturnSuccess(c, util.PublicSuccess, rid)
}

// @Summary 更新菜单
// @Tags Menu
// @Accept  x-www-form-urlencoded
// @Router /menu/update [put]
func (s *MenuController) UpdateMenu(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parent_id"))

	data, err := s.MenuService.UpdateMenu(id, &model.Menu{
		Name:     name,
		ParentId: parentId,
	})
	if err != nil {
		s.ReturnFail(c, util.PublicError, nil)
		return
	}
	s.ReturnSuccess(c, util.PublicSuccess, data)
}

// @Summary 获取菜单树
// @Tags Menu
// @Router /menu/tree [get]
func (s *MenuController) GetTree(c *gin.Context) {
	data, err := s.MenuService.GetAllMenu()
	if err != nil {
		s.ReturnFail(c, util.PublicError, nil)
		return
	}
	s.ReturnSuccess(c, util.PublicSuccess, data)
}
