package controllers

import (
	"fmt"
	"gin-base/internal/model"
	services "gin-base/internal/service"
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MangerController struct {
	BaseController
	ManagerService *services.ManagerService
}

// @Summary 管理员列表
// @Tags Manage
// @Produce  json
// @Param keyword query string false "keyword"
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Router /manager/list [get]
func (m *MangerController) GetManagerList(c *gin.Context) {
	// 初始化查询条件
	var query model.ManagerQuery

	// 接收分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 接收参数
	keyword := c.Query("keyword")

	// 设置查询条件
	query.Keyword = &keyword

	// 调用服务
	data, err := m.ManagerService.GetManagerList(query, page, pageSize)
	if err != nil {
		m.ReturnFail(c, util.PublicError, nil)
	}
	m.ReturnSuccess(c, util.PublicSuccess, data)
}

// @Summary 添加管理员
// @Tags Manage
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   name  formData  string  true  "用户名"
// @Param   password  formData  string  true  "密码"
// @Param   status formData  int  true  "状态" value 1 启用 0 禁用
// @Router /manager/add [post]
func (m *MangerController) AddManager(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "1"))

	if name == "" || password == "" {
		m.ReturnFail(c, util.PublicParamsIllegal, nil)
		return
	}
	var adModel model.AuthManager
	hashPass, _ := util.HashPassword(password)
	adModel.Name = name
	adModel.Password = hashPass
	adModel.Status = status
	adModel.UpIp = c.ClientIP()
	id, err := m.ManagerService.AddManager(&adModel)
	if err != nil {
		m.ReturnFail(c, util.PublicError, err)
		return
	}
	m.ReturnSuccess(c, util.PublicSuccess, id)
}

// @Summary 更新管理员
// @Tags Manage
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   id   path    int     true        "Manager ID"
// @Param   name  formData  string  true  "用户名"
// @Param   password  formData  string  true  "密码"
// @Param   status formData  int  true  "状态" value 1 启用 0 禁用
// @Router /manager/update [put]
func (m *MangerController) UpdateManager(c *gin.Context) {
	var upModel model.AuthManager
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		m.ReturnFail(c, util.PublicParamsIllegal, nil)
		return
	}
	name := c.PostForm("name")
	password := c.PostForm("password")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "1"))
	if name == "" || password == "" {
		m.ReturnFail(c, util.PublicParamsIllegal, nil)
		return
	}
	fmt.Println(upModel.ID, id)
	upPass, _ := util.HashPassword(password)
	upModel.Name = name
	upModel.Status = status
	upModel.Password = upPass
	res, err := m.ManagerService.UpdateManager(id, &upModel)
	if err != nil {
		m.ReturnFail(c, util.PublicError, err)
	}
	m.ReturnSuccess(c, util.PublicSuccess, res)
}

// @Summary 删除管理员
// @Tags Manage
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   id   path    int     true        "Manager ID"
// @Router /manager/delete [delete]
func (m *MangerController) DeleteManager(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		m.ReturnFail(c, util.PublicParamsIllegal, nil)
		return
	}
	err := m.ManagerService.DeleteManager(id)
	if err != nil {
		m.ReturnFail(c, util.PublicError, err)
	}
	m.ReturnSuccess(c, util.PublicSuccess, nil)
}
