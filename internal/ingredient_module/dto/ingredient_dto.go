package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"craeted_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
