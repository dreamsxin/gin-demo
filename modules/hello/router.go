package hello

import (
	"net/http"

	"gin-demo/logger"
	"gin-demo/modules/hello/models"
	"gin-demo/modules/hello/services"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	logger.Sugar().Debugln("hello router init")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		user := models.User{}
		services.SerUser.DB().First(&user)
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})
}
