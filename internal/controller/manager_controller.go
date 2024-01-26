package controllers

import (
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
