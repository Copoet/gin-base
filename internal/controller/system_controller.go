package controllers

import (
	"gin-base/internal/model"
	services "gin-base/internal/service"
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SystemController struct {
	BaseController
	SystemService *services.SystemService
}

// @Summary 获取系统配置列表
// @Tags System
// @Produce  json
// @Param sys_name query string false "sys_name"
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Router /system/list [get]
func (s *SystemController) GetSystemList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	SystemName := c.Query("sys_name")
	data, err := s.SystemService.GetSystemList(model.SystemQuery{
		SysName: &SystemName,
	}, page, pageSize)
	if err != nil {
		s.ReturnFail(c, util.PublicError, nil)
		return
	}
	s.ReturnSuccess(c, util.PublicSuccess, data)
}
