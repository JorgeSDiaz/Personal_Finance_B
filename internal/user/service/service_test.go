package service_test

import (
	"fmt"
	"testing"

	"github.com/JorgeSDiaz/Personal_Finance_B/internal/user/model"
	"github.com/JorgeSDiaz/Personal_Finance_B/internal/user/service"
	"github.com/stretchr/testify/assert"
)

func TestService_AllUsers(t *testing.T) {
	testCases := []struct {
		name             string
		mockRepository   *mockRepository
		expectedResponse []model.UserResponse
	}{
		{
			name:             "TestService - AllUsers - Empty Slice",
			mockRepository:   buildMockFindAllRepo(emptyUsers),
			expectedResponse: emptyUsersResponse,
		},
		{
			name:             "TestService - AllUsers - Completed Slice",
			mockRepository:   buildMockFindAllRepo(dummyUsers),
			expectedResponse: dummyUsersResponse,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			service := service.NewService(tc.mockRepository)
			response := service.AllUsers()

			for i := range tc.expectedResponse {
				assert.Equal(t, tc.expectedResponse[i], response[i])
			}
		})
	}
}

func TestService_Registry(t *testing.T) {
	testCases := []struct {
		name             string
		request          model.UserRegistryRequest
		mockRepository   *mockRepository
		expectedResponse model.UserResponse
		expectError      error
	}{
		{
			name:             "TestService - Registry - User Already Exist - Error",
			request:          dummyRegistryRequest,
			mockRepository:   buildMockSaveRepo(&dummyRegistryRequest, nil, fmt.Errorf("email already exist")),
			expectedResponse: model.UserResponse{},
			expectError:      fmt.Errorf("email already exist"),
		},
		{
			name:             "TestService - Registry - User Saved",
			request:          dummyRegistryRequest,
			mockRepository:   buildMockSaveRepo(&dummyRegistryRequest, &dummyUser, nil),
			expectedResponse: dummyUserResponse,
			expectError:      nil,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			service := service.NewService(tc.mockRepository)
			response, err := service.Registry(&tc.request)

			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectError, err)
		})
	}
}

func TestService_LogIn(t *testing.T) {
	testCases := []struct {
		name             string
		request          model.UserLogInRequest
		mockRepository   *mockRepository
		expectedResponse model.UserResponse
		expectError      error
	}{
		{
			name:             "TestService - LogIn - User Not Found - Error",
			request:          dummyBadLoginRequest,
			mockRepository:   buildMockGetUserByEmailRepo(&dummyBadLoginRequest, nil),
			expectedResponse: model.UserResponse{},
			expectError:      fmt.Errorf("user with email dummy@testing, not found"),
		},
		{
			name:             "TestService - LogIn - Wrong Password - Error",
			request:          dummyBadLoginRequest,
			mockRepository:   buildMockGetUserByEmailRepo(&dummyBadLoginRequest, &dummyUser),
			expectedResponse: model.UserResponse{},
			expectError:      fmt.Errorf("wrong password"),
		},
		{
			name:             "TestService - LogIn - Success",
			request:          dummyLoginRequest,
			mockRepository:   buildMockGetUserByEmailRepo(&dummyLoginRequest, &dummyUser),
			expectedResponse: dummyUserResponse,
			expectError:      nil,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			service := service.NewService(tc.mockRepository)
			response, err := service.LogIn(&tc.request)

			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectError, err)
		})
	}
}
