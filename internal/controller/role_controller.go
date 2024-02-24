package controllers

import (
	"gin-base/internal/model"
	services "gin-base/internal/service"
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RoleController struct {
	BaseController
	RoleService *services.RoleService
}

// @Summary 角色列表
// @Tags Role
// @Produce  json
// @Param name query string false "name"
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Router /role/list [get]
func (rc *RoleController) GetList(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	data, err := rc.RoleService.GetRoleList(model.QueryRole{}, page, pageSize)
	if err != nil {
		rc.ReturnFail(c, util.PublicError, err)
	}
	rc.ReturnSuccess(c, util.PublicSuccess, data)
}

// @Summary 获取单个角色信息
// @Tags Role
// @Produce  json
// @Param id path int true "id"
// @Router /role/info/{id} [get]
func (rc *RoleController) GetRoleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := rc.RoleService.GetRoleInfo(
		model.QueryRole{ID: &id})
	if err != nil {
		rc.ReturnFail(c, util.PublicError, err)
	}
	rc.ReturnSuccess(c, util.PublicSuccess, data)

}

// @Summary 添加角色
// @Tags Role
// @Accept x-www-form-urlencoded
// @Produce json
// @Param name formData string true "角色名称"
// @Router /role/add [post]
func (rc *RoleController) AddRole(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		rc.ReturnFail(c, util.PublicParamsNull, nil)
	}
	status, err := strconv.Atoi(c.PostForm("status"))
	if err != nil {
		rc.ReturnFail(c, util.PublicParamsNull, nil)
	}
	data, err := rc.RoleService.AddRole(&model.Role{
		Name:   name,
		Status: status,
	})
	if err != nil {
		rc.ReturnFail(c, util.PublicError, err)
	}
	rc.ReturnSuccess(c, util.PublicSuccess, data)
}
