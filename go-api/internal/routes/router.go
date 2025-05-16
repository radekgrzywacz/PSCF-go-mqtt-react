package routes

import (
	"go-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/test", handler.Test)
	}

	return r
}
