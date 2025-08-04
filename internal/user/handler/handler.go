package handler

import (
	"net/http"

	"github.com/JorgeSDiaz/Personal_Finance_B/internal/user/model"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service UserService
}

func NewHandler(s UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Registry(c *gin.Context) {
	var request model.UserRegistryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := h.service.Registry(&request)
	if err == nil {
		c.JSON(http.StatusOK, userResponse)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func (h *UserHandler) LogIn(c *gin.Context) {
	var request model.UserLogInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := h.service.LogIn(&request)
	if err == nil {
		c.JSON(http.StatusOK, userResponse)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func (h *UserHandler) Users(c *gin.Context) {
	usersResponse := h.service.AllUsers()
	c.JSON(http.StatusOK, usersResponse)
}
