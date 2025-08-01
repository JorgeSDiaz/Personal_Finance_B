package handler_test

import (
	"errors"

	"github.com/JorgeSDiaz/Personal_Finance_B/internal/user/model"
)

// Errors
var (
	errAlready = errors.New("email already exist")
	errInvalid = errors.New("invalid credentials")
)

// UserRegistryRequest
var (
	dummyRawRegistryRequest = `{"fullName":"dummy testing de leon","email":"deLeon@testing.com","password":"dummy123"}`

	dummyUserRegistryRequest = model.UserRegistryRequest{
		FullName: "dummy testing de leon",
		Email:    "deLeon@testing.com",
		Password: "dummy123",
	}
)

// UserLogInRequest
var (
	dummyRawLogInRequest = `{"email":"deLeon@testing.com","password":"dummy123"}`

	dummyUserLogInRequest = model.UserLogInRequest{
		Email:    "deLeon@testing.com",
		Password: "dummy123",
	}
)

// UserResponse
var (
	dummyUserResponse = model.UserResponse{
		ID:       1,
		FullName: "dummy testing de leon",
		Email:    "deLeon@testing.com",
	}

	dummyUsersResponse = []model.UserResponse{
		{
			ID:       1,
			FullName: "dummy testing de leon",
			Email:    "deLeon@testing.com",
		},
		{
			ID:       2,
			FullName: "dummy testing aguilares",
			Email:    "aguilares@testing.com",
		},
	}
)
