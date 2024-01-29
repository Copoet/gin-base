package controllers

import (
	"gin-base/internal/model"
	services "gin-base/internal/service"
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	BaseController
	ManagerService *services.ManagerService
}

// @Summary 登陆
// @Accept  json
// @Produce  json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Param c body object true "Context"
// @Router /auth/login [post]
func (a *AuthController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 1.参数校验
	if username == "" || password == "" {
		a.ReturnFail(c, util.PublicParamsIllegal, nil)
		return
	}
	// 2.根据用户名和密码去数据库查询
	var queryWhere model.ManagerQuery
	queryWhere.Name = &username
	authInfo, err := a.ManagerService.GetManagerInfo(queryWhere)
	if err != nil {
		a.ReturnFail(c, util.PublicError, err)
		return
	}
	if authInfo == nil {
		a.ReturnSuccess(c, util.PublicAuthError, nil)
		return
	}

	if !util.CheckPasswordHash(password, authInfo.Password) {
		a.ReturnSuccess(c, util.PublicAuthError, nil)
		return
	}
	// 3.生成token
	token, err := util.GenerateToken(authInfo.Name, authInfo.Password)
	if err != nil {
		a.ReturnFail(c, util.PublicError, err)
		return
	}
	a.ReturnSuccess(c, util.PublicSuccess, token)
	return
}

// @Summary 退出登陆
// @param c
// @return
// @Router /auth/logout [post]
func (a *AuthController) Logout(c *gin.Context) {
	// 获取token
	tokenString := c.GetHeader("Authorization")
	// 校验token
	if !util.ValidTokenExpress(tokenString) {
		a.ReturnFail(c, util.PublicAuthError, nil)
		return
	}
	// 删除token
	util.DeleteToken(tokenString)
	a.ReturnSuccess(c, util.PublicSuccess, nil)

}
