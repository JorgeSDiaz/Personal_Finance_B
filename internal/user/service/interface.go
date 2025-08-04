package service

import "github.com/JorgeSDiaz/Personal_Finance_B/internal/user/model"

type Repository interface {
	FindAll() []*model.User
	Save(userRequest *model.UserRegistryRequest) (*model.User, error)
	GetUserByEmail(userRequest *model.UserLogInRequest) *model.User
}
