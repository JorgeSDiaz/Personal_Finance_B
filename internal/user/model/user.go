package model

import (
	"time"
)

type UserRegistryRequest struct {
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserLogInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type User struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}
