package handler

import models "github.com/JorgeSDiaz/Personal_Finance_B/internal/health/model"

type HealthService interface {
	Check() models.HealthResponse
}
