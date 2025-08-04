package handler_test

import (
	"github.com/JorgeSDiaz/Personal_Finance_B/internal/user/model"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (m *mockService) Registry(userRequest *model.UserRegistryRequest) (model.UserResponse, error) {
	arg := m.Called(userRequest)

	return arg.Get(0).(model.UserResponse), arg.Error(1)
}

func buildMockRegistryService(userRequest *model.UserRegistryRequest, expectedUserResponse model.UserResponse, expectedError error) *mockService {
	m := new(mockService)

	m.On("Registry", userRequest).Return(expectedUserResponse, expectedError)
	return m
}

func (m *mockService) LogIn(userRequest *model.UserLogInRequest) (model.UserResponse, error) {
	arg := m.Called(userRequest)

	return arg.Get(0).(model.UserResponse), arg.Error(1)
}

func buildMockLogInService(userRequest *model.UserLogInRequest, expectedUserResponse model.UserResponse, expectedError error) *mockService {
	m := new(mockService)

	m.On("LogIn", userRequest).Return(expectedUserResponse, expectedError)
	return m
}

func (m *mockService) AllUsers() []model.UserResponse {
	arg := m.Called()

	return arg.Get(0).([]model.UserResponse)
}

func buildMockAllUsersService(expectedUsers []model.UserResponse) *mockService {
	m := new(mockService)

	m.On("AllUsers").Return(expectedUsers)
	return m
}
