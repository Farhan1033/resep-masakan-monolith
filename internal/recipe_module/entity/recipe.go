package recipeentity

import (
	"time"

	"github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/entity"
	categoryentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Recipe struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title          string    `gorm:"type:varchar(255);not null" json:"title"`
	Description    string    `gorm:"type:text;not null" json:"description"`
	DifficultLevel string    `gorm:"type:varchar(6);enum('EASY','MEDIUM','HARD')" json:"difficult_level"`
	PrepTime       int       `gorm:"type:integer;not null" json:"prep_time"`
	CookTime       int       `gorm:"type:integer;not null" json:"cook_time"`
	TotalTime      int       `gorm:"type:integer;not null" json:"total_time"`
	Servings       int       `gorm:"type:integer;not null" json:"servings"`
	OriginRegion   string    `gorm:"type:varchar(100);not null" json:"origin_region"`
	ImageUrl       string    `gorm:"type:text" json:"image_url"`
	CategoryId     uuid.UUID `gorm:"type:uuid;not null" json:"category_id"`
	CreatedBy      uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	IsActive       bool      `gorm:"type:boolean;default:true" json:"is_active"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updated_at"`
	// Relation Table
	User     entity.User             `gorm:"foreignKey:CreatedBy;references:ID" json:"users"`
	Category categoryentity.Category `gorm:"foreignKey:CategoryId;references:ID" json:"category"`
}

func (r *Recipe) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}

	return
}

func (Recipe) TableName() string {
	return "recipe"
}
