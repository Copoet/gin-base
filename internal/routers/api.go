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
	articleController := new(controllers.ArticleController)
	categoryController := new(controllers.CategoryController)
	//测试中间价
	router.Use(middleware.LoggerMiddleware())
	// login & logout
	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.GET("/logout", authController.Logout)
	}

	//用户
	users := router.Group("/users", middleware.JwtMiddleware())
	{
		users.GET("/list", userController.GetList)
	}
	// 管理员
	manager := router.Group("/manager")
	{
		manager.GET("/list", managerController.GetManagerList)
		manager.POST("/add", managerController.AddManager)
		manager.PUT("/update/:id", managerController.UpdateManager)
		manager.DELETE("/delete/:id", managerController.DeleteManager)
	}
	// 文章
	article := router.Group("/article")
	{
		article.GET("/list", articleController.GetList)
		article.POST("/add", articleController.AddArticle)
		article.PUT("/update/:id", articleController.UpdateArticle)
		article.DELETE("/delete/:id", articleController.DeleteArticle)
	}
	//分类
	category := router.Group("/category")
	{
		category.GET("/list", categoryController.GetList)
		category.POST("/add", categoryController.AddCategory)
		category.PUT("/update/:id", categoryController.UpdateCategory)
		category.DELETE("/delete/:id", categoryController.DeleteCategory)
	}
}
