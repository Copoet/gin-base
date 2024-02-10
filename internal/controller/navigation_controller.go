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
// @Router /navigation/list [get]
func (n *NavigationController) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
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
