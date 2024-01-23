package routers

import (
	controllers "gin-base/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	userController := new(controllers.UserController)
	//router.POST("/users", userController.Create)
	router.GET("/users/list", userController.GetList)
}
