package controllers

import (
	"fmt"
	services "gin-base/internal/service"
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
	BaseController
	UserService *services.UserService
}

// @Summary Get User List
// @Tags User
// @Produce  json
// @Param name query string false "name"
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Router /users/list [get]
func (u *UserController) GetList(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	name := c.Query("name")
	fmt.Println(util.HashPassword("123456"))
	data, err := u.UserService.GetUsers(page, pageSize, name)
	if err != nil {
		u.ReturnFail(c, util.PublicError, nil)
	}
	u.ReturnSuccess(c, util.PublicSuccess, data)
}
