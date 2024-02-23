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
