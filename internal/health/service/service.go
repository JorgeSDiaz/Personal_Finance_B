package service

import models "github.com/JorgeSDiaz/Personal_Finance_B/internal/health/model"

type HealthService struct{}

func NewService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) Check() *models.HealthResponse {
	return &models.HealthResponse{
		Response: "Ok!",
	}
}
