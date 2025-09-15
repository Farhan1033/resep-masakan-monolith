package dto

import "github.com/google/uuid"

type CreateRequest struct {
	Name     string `json:"name" validate:"required,min=4,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=30"`
}

type UpdateRequest struct {
	Name     string `json:"name,omitempty" validate:"omitempty,min=4,max=30"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Password string `json:"password,omitempty" validate:"omitempty,min=6,max=30"`
}

type CreateAuthResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
