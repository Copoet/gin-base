package controllers

import (
	"gin-base/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

type ResponseData struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

// 统一返回成功格式｜ReturnSuccess is a method of the BaseController struct that formats and returns a successful response.
func (bc *BaseController) ReturnSuccess(c *gin.Context, code int, data interface{}) {
	msg := util.GetMsg(code)
	c.JSON(http.StatusOK, ResponseData{
		Code:   code,
		Msg:    msg,
		Status: true,
		Data:   data,
	})
}

// 统一返回失败格式｜ReturnFail formats and returns a failure response.
func (bc *BaseController) ReturnFail(c *gin.Context, code int, data interface{}) {
	msg := util.GetMsg(code)
	c.JSON(http.StatusOK, ResponseData{
		Code:   code,
		Msg:    msg,
		Status: false,
		Data:   data,
	})
}
