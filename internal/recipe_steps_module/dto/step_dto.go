package stepdto

import (
	"time"

	"github.com/google/uuid"
)

type CreateRequest struct {
	RecipeId    uuid.UUID `json:"recipe_id" validate:"required"`
	Instruction string    `json:"instruction" validate:"required"`
}

type CreateResponse struct {
	ID          uint      `json:"id"`
	RecipeId    uuid.UUID `json:"recipe_id"`
	StepNumber  int       `json:"step_number"`
	Instruction string    `json:"instruction"`
	CreatedAt   time.Time `json:"created_at"`
}

type RecipeStepResponse struct {
	ID          uint      `json:"id"`
	RecipeId    uuid.UUID `json:"recipe_id"`
	StepNumber  int       `json:"step_number"`
	Instruction string    `json:"instruction"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateRequest struct {
	Instruction string `json:"instruction,omitempty" validate:"omitempty"`
}
