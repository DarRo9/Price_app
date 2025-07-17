package delivery

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *ImageRecognitionHandler) *gin.Engine {
	r := gin.Default()

	r.StaticFile("/", "./web/index.html")
	r.Static("/static", "./web/static")

	api := r.Group("/api/v1")
	{
		api.POST("/upload", handler.UploadImage)
		api.GET("/health", handler.HealthCheck)
	}

	return r
}
