package dto

import "github.com/google/uuid"

type CreateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	UserName string `json:"user_name" validate:"required,min=4"`
	FullName string `json:"full_name" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=6,max=30"`
}

type UpdateRequest struct {
	Email     string `json:"email,omitempty" validate:"omitempty,email"`
	FullName  string `json:"full_name,omitempty" validate:"omitempty,min=4,max=30"`
	AvatarUrl string `json:"avatar_url,omitempty" validate:"omitempty;url"`
	Bio       string `json:"bio,omitempty" validate:"omitempty,"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=30"`
}

type CreateAuthResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	FullName  string    `json:"full_name"`
	CreatedAt string    `json:"created_at"`
}

type LoginResponse struct {
	Data  *LoginDataResponse `json:"data"`
	Token any                `json:"token"`
}

type LoginDataResponse struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	UserName string    `json:"user_name"`
	FullName string    `json:"full_name"`
}
