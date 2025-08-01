package api

import (
	healthHandler "github.com/JorgeSDiaz/Personal_Finance_B/internal/health/handler"
	healthService "github.com/JorgeSDiaz/Personal_Finance_B/internal/health/service"

	userHandler "github.com/JorgeSDiaz/Personal_Finance_B/internal/user/handler"
	userRepository "github.com/JorgeSDiaz/Personal_Finance_B/internal/user/repository"
	userService "github.com/JorgeSDiaz/Personal_Finance_B/internal/user/service"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetUpRoutes() *gin.Engine {
	if router == nil {
		router = gin.Default()
	}

	router = SetUpHealthRoutes(router)
	router = SetUpUserRoutes(router)

	return router
}

func SetUpHealthRoutes(router *gin.Engine) *gin.Engine {
	service := healthService.NewService()
	handler := healthHandler.NewHandler(service)

	health := router.Group("/health")
	{
		health.GET("", handler.Check)
	}

	return router
}

func SetUpUserRoutes(router *gin.Engine) *gin.Engine {
	repository := userRepository.NewInMemoryRepository()
	service := userService.NewService(repository)
	handler := userHandler.NewHandler(service)

	user := router.Group("/users")
	{
		user.GET("", handler.Users)
		user.POST("/login", handler.LogIn)
		user.POST("/registry", handler.Registry)
	}

	return router
}
