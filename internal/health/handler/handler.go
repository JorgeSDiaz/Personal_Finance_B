package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	service HealthService
}

func NewHandler(s HealthService) *HealthHandler {
	return &HealthHandler{service: s}
}

func (h *HealthHandler) Check(c *gin.Context) {
	healthResponse := h.service.Check()
	c.JSON(http.StatusOK, healthResponse)
}
