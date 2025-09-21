package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateRequest struct {
	CategoryId     uuid.UUID `json:"category_id" validaet:"required"`
	Title          string    `json:"title" validate:"required"`
	Description    string    `json:"description" validate:"required"`
	DifficultLevel string    `json:"difficult_level" validate:"required,oneof=EASY MEDIUM HARD"`
	PrepTime       int       `json:"prep_time" validate:"required"`
	CookTime       int       `json:"cook_time" validate:"required"`
	TotalTime      int       `json:"total_time" validate:"required"`
	Servings       int       `json:"servings" validate:"required"`
	OriginRegion   string    `json:"origin_region" validate:"required"`
	ImageUrl       string    `json:"image_url" validate:"required,url"`
}

type CreateResponse struct {
	ID             uuid.UUID  `json:"id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	DifficultLevel string     `json:"difficult_level"`
	PrepTime       int        `json:"prep_time"`
	CookTime       int        `json:"cook_time"`
	TotalTime      int        `json:"total_time"`
	Servings       int        `json:"servings"`
	OriginRegion   string     `json:"origin_region"`
	ImageUrl       string     `json:"image_url"`
	CreatedAt      time.Time  `json:"created_at"`
	Category       *Category  `json:"category"`
	CreatedBy      *CreatedBy `json:"created_by"`
}

type CreatedBy struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
}

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetResponse struct {
	ID             uuid.UUID  `json:"id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	DifficultLevel string     `json:"difficult_level"`
	PrepTime       int        `json:"prep_time"`
	CookTime       int        `json:"cook_time"`
	TotalTime      int        `json:"total_time"`
	Servings       int        `json:"servings"`
	OriginRegion   string     `json:"origin_region"`
	ImageUrl       string     `json:"image_url"`
	Category       *Category  `json:"category"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdateAt       time.Time  `json:"updated_at"`
	CreatedBy      *CreatedBy `json:"created_by"`
}

type PaginatedResponse struct {
	Data        []*GetResponse `json:"recipe"`
	CurrentPage int            `json:"current_page"`
	PerPage     int            `json:"per_page"`
	Total       int            `json:"total"`
	TotalPages  int            `json:"total_pages"`
	HasNext     bool           `json:"has_next"`
	HasPrev     bool           `json:"has_prev"`
}

type UpdateRequest struct {
	CategoryId     uuid.UUID `json:"category_id" validate:"omitempty"`
	Title          string    `json:"title,omitempty" validate:"omitempty"`
	Description    string    `json:"description,omitempty" validate:"omitempty"`
	DifficultLevel string    `json:"difficult_level,omitempty" validate:"omitempty,oneof=EASY MEDIUM HARD"`
	PrepTime       int       `json:"prep_time,omitempty" validate:"omitempty,gte=0"`
	CookTime       int       `json:"cook_time,omitempty" validate:"omitempty,gte=0"`
	TotalTime      int       `json:"total_time,omitempty" validate:"omitempty,gte=0"`
	Servings       int       `json:"servings,omitempty" validate:"omitempty,gte=0"`
	OriginRegion   string    `json:"origin_region,omitempty" validate:"omitempty"`
	ImageUrl       string    `json:"image_url,omitempty" validate:"omitempty,url"`
}

type UpdateResponse struct {
	ID             uuid.UUID `json:"id"`
	CategoryId     uuid.UUID `json:"category_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	DifficultLevel string    `json:"difficult_level"`
	PrepTime       int       `json:"prep_time"`
	CookTime       int       `json:"cook_time"`
	TotalTime      int       `json:"total_time"`
	Servings       int       `json:"servings"`
	OriginRegion   string    `json:"origin_region"`
	ImageUrl       string    `json:"image_url"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type RecipeWithRelations struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	DifficultLevel string    `json:"difficult_level"`
	PrepTime       int       `json:"prep_time"`
	CookTime       int       `json:"cook_time"`
	TotalTime      int       `json:"total_time"`
	Servings       int       `json:"servings"`
	OriginRegion   string    `json:"origin_region"`
	ImageUrl       string    `json:"image_url"`
	IsActive       bool      `json:"is_active"`
	UserID         uuid.UUID `json:"user_id"`
	UserName       string    `json:"user_name"`
	CategoryID     uuid.UUID `json:"category_id"`
	CategoryName   string    `json:"category_name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
