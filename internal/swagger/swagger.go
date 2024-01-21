package swagger

import (
	"gin-base/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupSwagger(r *gin.Engine) {
	url := ginSwagger.URL(config.Config.App.PrefixUrl + "/swagger/doc.json") // Swagger JSON的URL
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
