package stepentity

import (
	"time"

	recipeentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/entity"
	"github.com/google/uuid"
)

type RecipeStep struct {
	ID          uint      `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	RecipeId    uuid.UUID `gorm:"type:uuid;not null" json:"recipe_id"`
	StepNumber  int       `gorm:"type:int;not null" json:"step_number"`
	Instruction string    `gorm:"type:text;not null" json:"instruction"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relation
	Recipe recipeentity.Recipe `gorm:"foreignKey:RecipeId;references:ID" json:"recipe"`
}

func (RecipeStep) TableName() string {
	return "recipe_steps"
}
