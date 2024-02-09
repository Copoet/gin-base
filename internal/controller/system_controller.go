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

// @Summary 添加系统配置
// @Tags System
// @Produce  json
// @Param sys_name formData string true "sys_name"
// @Param sys_value formData string true "sys_value"
// @Param sys_type formData string true "sys_type"
// @Router /system/add [post]
func (s *SystemController) AddSystem(c *gin.Context) {
	sysName := c.PostForm("sys_name")
	sysValue := c.PostForm("sys_value")
	sysType := c.PostForm("sys_type")
	if sysName == "" || sysValue == "" || sysType == "" {
		s.ReturnFail(c, util.PublicError, nil)
		return
	}
	system := &model.System{
		SysName:    sysName,
		SysValue:   sysValue,
		SysExplain: sysName,
		SysType:    sysType,
		IsDelete:   0,
		Status:     1,
	}
	if _, err := s.SystemService.AddSystem(system); err != nil {
		s.ReturnFail(c, util.PublicError, err)
		return
	}
	s.ReturnSuccess(c, util.PublicSuccess, nil)

}

// @Summary 更新系统配置
// @Tags System
// @Produce  json
// @Param id formData int true "id"
// @Param sys_name formData string true "sys_name"
// @Param sys_value formData string true "sys_value"
// @Param sys_type formData string true "sys_type"
// @Router /system/update [put]
func (s *SystemController) UpdateSystem(c *gin.Context) {
	sysName := c.PostForm("sys_name")
	sysValue := c.PostForm("sys_value")
	sysType := c.PostForm("sys_type")
	status, _ := strconv.Atoi(c.PostForm("status"))
	id, _ := strconv.Atoi(c.PostForm("id"))
	if sysName == "" || sysValue == "" || sysType == "" {
		s.ReturnFail(c, util.PublicError, nil)
		return
	}
	system := &model.System{
		SysName:    sysName,
		SysValue:   sysValue,
		SysExplain: sysName,
		SysType:    sysType,
		IsDelete:   0,
		Status:     status,
	}
	if _, err := s.SystemService.UpdateSystem(id, system); err != nil {
		s.ReturnFail(c, util.PublicError, err)
		return
	}
	s.ReturnSuccess(c, util.PublicSuccess, nil)

}

// @Summary 删除系统配置
// @Tags System
// @Produce  json
// @Param id formData int true "id"
// @Router /system/delete [delete]
func (s *SystemController) DeleteSystem(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	data, err := s.SystemService.DeleteSystem(id)
	if err != nil {
		s.ReturnFail(c, util.PublicError, nil)
		return
	}
	s.ReturnSuccess(c, util.PublicSuccess, data)
}
