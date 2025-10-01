package detailrecipeentity

import (
	"time"

	ingrediententity "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/entity"
	recipeentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DetailRecipeEntity struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	RecipeID     uuid.UUID `gorm:"type:uuid;not null" json:"recipe_id"`
	IngredientID uuid.UUID `gorm:"type:uuid;not null" json:"ingredient_id"`
	Quantity     float64   `gorm:"type:decimal(10,2);not null" json:"quantity"`
	Unit         string    `gorm:"type:varchar(50);not null" json:"unit"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updated_at"`
	CreatedBy    string    `gorm:"varchar(100)" json:"created_by"`

	// Relations
	Recipe     recipeentity.Recipe         `gorm:"foreignKey:RecipeID;references:ID" json:"recipe"`
	Ingredient ingrediententity.Ingredient `gorm:"foreignKey:IngredientID;references:ID" json:"ingredient"`
}

func (ri *DetailRecipeEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if ri.ID == uuid.Nil {
		ri.ID = uuid.New()
	}

	return
}

func (DetailRecipeEntity) TableName() string {
	return "detail_recipes"
}
