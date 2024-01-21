package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
	}))
}
