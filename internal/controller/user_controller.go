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

func (u *UserController) GetList(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.Query("name")

	users, count, err := u.UserService.GetUsers(page, pageSize, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      users,
		"total":     count,
		"page":      page,
		"page_size": pageSize,
	})
}
