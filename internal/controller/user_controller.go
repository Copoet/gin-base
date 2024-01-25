package controllers

import (
	services "gin-base/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService *services.UserService
}

// GetList 方法用于获取用户列表
// @Summary 获取用户列表
// @Description 根据分页和名字参数获取用户列表
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "分页大小"
// @Param name query string false "用户名"
// @Success 200 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [get]
func (u *UserController) GetList(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.Query("name")

	data, err := u.UserService.GetUsers(page, pageSize, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
