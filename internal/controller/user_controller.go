package controllers

import (
	"gin-base/internal/database"
	"gin-base/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type UserController struct {
	DB *gorm.DB //
}

func (u *UserController) GetList(c *gin.Context) {
	var users []model.User
	var count int64

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.Query("name")

	offset := (page - 1) * pageSize

	query := database.DB.Model(&model.User{})
	if name != "" {
		query = query.Where("username LIKE ?", "%"+name+"%")
	}

	query.Count(&count).Limit(pageSize).Offset(offset).Find(&users)

	c.JSON(200, gin.H{
		"data":      users,
		"total":     count,
		"page":      page,
		"page_size": pageSize,
	})
}
