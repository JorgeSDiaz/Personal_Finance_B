package service_test

import (
	"github.com/JorgeSDiaz/Personal_Finance_B/internal/user/model"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) Save(userRequest *model.UserRegistryRequest) (*model.User, error) {
	args := m.Called(userRequest)
	return args.Get(0).(*model.User), args.Error(1)
}

func buildMockSaveRepo(expectedRequest *model.UserRegistryRequest, expectedUser *model.User, expectedError error) *mockRepository {
	m := new(mockRepository)

	m.On("Save", expectedRequest).Return(expectedUser, expectedError)
	return m
}

func (m *mockRepository) GetUserByEmail(userRequest *model.UserLogInRequest) *model.User {
	args := m.Called(userRequest)
	return args.Get(0).(*model.User)
}

func buildMockGetUserByEmailRepo(expectedRequest *model.UserLogInRequest, expectedUser *model.User) *mockRepository {
	m := new(mockRepository)

	m.On("GetUserByEmail", expectedRequest).Return(expectedUser)
	return m
}

func (m *mockRepository) FindAll() []*model.User {
	args := m.Called()
	return args.Get(0).([]*model.User)
}

func buildMockFindAllRepo(expectedUsers []*model.User) *mockRepository {
	m := new(mockRepository)

	m.On("FindAll").Return(expectedUsers)
	return m
}
