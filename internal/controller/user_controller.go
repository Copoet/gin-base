package controllers

import (
	"gin-base/internal/model"
	"gin-base/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u *UserController) Create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 数据验证
	if err := services.CheckUser(&user); err != nil {
		c.JSON(StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	c.JSON(StatusOK, user)
}
