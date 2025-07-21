package health

type HealthService struct{}

func NewService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) Check() *HealthResponse {
	return &HealthResponse{
		Response: "Ok!",
	}
}
