package controllers

import (
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	BaseController
}

func (a *AuthController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 1.参数校验
	if username == "" || password == "" {
		a.ReturnFail(c, util.PublicParamsIllegal, nil)
		return
	}

	// 2.根据用户名和密码去数据库查询
	// 3.返回响应
	// 4.登录成功后,生成token
	// 5.返回token
	// 6.返回响应

}
