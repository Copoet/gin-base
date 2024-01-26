package routers

import (
	controllers "gin-base/internal/controller"
	"gin-base/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	userController := new(controllers.UserController)
	managerController := new(controllers.MangerController)

	//router.POST("/users", userController.Create)
	//测试中间价
	router.Use(middleware.LoggerMiddleware())

	router.GET("/users/list", userController.GetList)

	router.GET("/manager/list", managerController.GetManagerList)
}
