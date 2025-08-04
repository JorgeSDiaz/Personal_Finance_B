package service

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/JorgeSDiaz/Personal_Finance_B/internal/user/model"
)

type UserService struct {
	repo Repository
}

func NewService(r Repository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Registry(userRequest *model.UserRegistryRequest) (model.UserResponse, error) {
	userSaved, err := s.repo.Save(userRequest)
	if err != nil {
		return model.UserResponse{}, err
	}

	return model.UserResponse{
		ID:       userSaved.ID,
		FullName: userSaved.FullName,
		Email:    userSaved.Email,
	}, nil
}

func (s *UserService) LogIn(userRequest *model.UserLogInRequest) (model.UserResponse, error) {
	user := s.repo.GetUserByEmail(userRequest)
	if user == nil {
		return model.UserResponse{}, fmt.Errorf("user with email %s, not found", userRequest.Email)
	}

	encodedRequestPassword := base64.StdEncoding.EncodeToString([]byte(userRequest.Password))
	if !strings.EqualFold(user.Password, encodedRequestPassword) {
		return model.UserResponse{}, fmt.Errorf("wrong password")
	}

	return model.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}, nil
}

func (s *UserService) AllUsers() []model.UserResponse {
	usersInMemory := s.repo.FindAll()

	var users = make([]model.UserResponse, 0)
	for _, user := range usersInMemory {
		users = append(users, model.UserResponse{
			ID:       user.ID,
			FullName: user.FullName,
			Email:    user.Email,
		})
	}

	return users
}
