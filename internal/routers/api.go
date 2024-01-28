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
	// login & logout
	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.GET("/logout", authController.Logout)
	}

	//users group
	users := router.Group("/users")
	{
		users.GET("/list", userController.GetList)
	}
	// manager group
	manager := router.Group("/manager")
	{
		manager.GET("/list", managerController.GetManagerList)
	}

}
