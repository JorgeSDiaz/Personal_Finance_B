package service_test

import (
	"time"

	"github.com/JorgeSDiaz/Personal_Finance_B/internal/user/model"
)

// TestData - Request
var (
	dummyRegistryRequest = model.UserRegistryRequest{
		FullName: "dummy testing registry service",
		Email:    "dummy@testing",
		Password: "dummy123",
	}

	dummyLoginRequest = model.UserLogInRequest{
		Email:    "dummy@testing",
		Password: "dummy123",
	}

	dummyBadLoginRequest = model.UserLogInRequest{
		Email:    "dummy@testing",
		Password: "dummyBad123",
	}
)

// TestData - UserResponse
var (
	emptyUsersResponse = make([]model.UserResponse, 0)
	dummyUsersResponse = []model.UserResponse{
		{
			ID:       1,
			FullName: "testing 1",
			Email:    "testing@testing.com",
		},
		{
			ID:       2,
			FullName: "testing 2",
			Email:    "testing2@testing.com",
		},
	}

	dummyUserResponse = model.UserResponse{
		ID:       1,
		FullName: "dummy testing registry service",
		Email:    "dummy@testing",
	}
)

// TestData - User
var (
	emptyUsers = make([]*model.User, 0)
	dummyUsers = []*model.User{
		{
			ID:        1,
			FullName:  "testing 1",
			Email:     "testing@testing.com",
			Password:  "ZHVtbXkxMjM=",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			FullName:  "testing 2",
			Email:     "testing2@testing.com",
			Password:  "ZHVtbXkxMjM=",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	dummyUser = model.User{
		ID:        1,
		FullName:  "dummy testing registry service",
		Email:     "dummy@testing",
		Password:  "ZHVtbXkxMjM=",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)
