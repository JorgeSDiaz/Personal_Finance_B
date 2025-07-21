package api

import (
	"github.com/JorgeSDiaz/Personal_Finance_B/internal/health"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetUpRoutes() *gin.Engine {
	if router == nil {
		router = gin.Default()
	}

	router = SetUpHealthRoutes(router)

	return router
}

func SetUpHealthRoutes(router *gin.Engine) *gin.Engine {
	service := health.NewService()
	handler := health.NewHandler(service)

	health := router.Group("/health")
	{
		health.GET("", handler.Check)
	}

	return router
}
