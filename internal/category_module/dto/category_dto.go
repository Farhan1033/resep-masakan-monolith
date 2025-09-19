package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=5"`
}

type CategoryResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
