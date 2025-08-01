package handler

import "github.com/JorgeSDiaz/Personal_Finance_B/internal/user/model"

type UserService interface {
	Registry(userRequest *model.UserRegistryRequest) (model.UserResponse, error)
	LogIn(userRequest *model.UserLogInRequest) (model.UserResponse, error)
	AllUsers() []model.UserResponse
}
