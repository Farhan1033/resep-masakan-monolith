package dto

import (
	"time"

	"github.com/google/uuid"
)

type RecipeWithIngredients struct {
	ID                uuid.UUID `json:"id"`
	RecipeID          uuid.UUID `json:"recipe_id"`
	RecipeTitle       string    `json:"title"`
	RecipeDescription string    `json:"description"`
	DifficultLevel    string    `json:"difficult_level"`
	PrepTime          int       `json:"prep_time"`
	CookTime          int       `json:"cook_time"`
	TotalTime         int       `json:"total_time"`
	Servings          int       `json:"servings"`
	OriginRegion      string    `json:"origin_region"`
	ImageURL          string    `json:"image_url"`
	IsActive          bool      `json:"is_active"`
	IngredientID      uuid.UUID `json:"ingredient_id"`
	IngredientName    string    `json:"ingredient_name"`
	Quantity          float64   `json:"quantity"`
	Unit              string    `json:"unit"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	CreatedBy         string    `json:"created_by"`
}

type CreateDetailRecipeRequest struct {
	RecipeID     uuid.UUID `json:"recipe_id" validate:"required"`
	IngredientID uuid.UUID `json:"ingredient_id" validate:"required"`
	Quantity     float64   `json:"quantity" validate:"required,gt=0"`
	Unit         string    `json:"unit" validate:"required"`
}

type UpdateDetailRecipeRequest struct {
	RecipeID     uuid.UUID `json:"recipe_id,omitempty" validate:"required"`
	IngredientID uuid.UUID `json:"ingredient_id,omitempty" validate:"required"`
	Quantity     float64   `json:"quantity,omitempty" validate:"required,gt=0"`
	Unit         string    `json:"unit,omitempty" validate:"required"`
}

type DetailRecipeResponse struct {
	ID           uuid.UUID `json:"id"`
	RecipeID     uuid.UUID `json:"recipe_id"`
	IngredientID uuid.UUID `json:"ingredient_id"`
	Quantity     float64   `json:"quantity"`
	Unit         string    `json:"unit"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
}

type IngredientDetail struct {
	IngredientID   uuid.UUID `json:"ingredient_id"`
	IngredientName string    `json:"ingredient_name"`
	Quantity       float64   `json:"quantity"`
	Unit           string    `json:"unit"`
}

type RecipeWithIngredientsResponse struct {
	ID                uuid.UUID          `json:"id"`
	RecipeID          uuid.UUID          `json:"recipe_id"`
	RecipeTitle       string             `json:"title"`
	RecipeDescription string             `json:"description"`
	DifficultLevel    string             `json:"difficult_level"`
	PrepTime          int                `json:"prep_time"`
	CookTime          int                `json:"cook_time"`
	TotalTime         int                `json:"total_time"`
	Servings          int                `json:"servings"`
	OriginRegion      string             `json:"origin_region"`
	ImageURL          string             `json:"image_url"`
	IsActive          bool               `json:"is_active"`
	Ingredients       []IngredientDetail `json:"ingredients"`
	CreatedBy         string             `json:"created_by"`
}
