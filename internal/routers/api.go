package routers

import (
	controllers "gin-base/internal/controller"
	"gin-base/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	userController := new(controllers.UserController)
	managerController := new(controllers.MangerController)
	authController := new(controllers.AuthController)

	//测试中间价
	router.Use(middleware.LoggerMiddleware())
	// login
	router.Group("/auth")
	{
		router.POST("/login", authController.Login)
		router.GET("/logout", authController.Logout)
	}

	//users group
	router.Group("/users")
	{
		router.GET("/list", userController.GetList)
	}
	// manager group
	router.Group("/manager")
	{
		router.GET("/manager/list", managerController.GetManagerList)
	}

}
