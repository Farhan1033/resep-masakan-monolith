package ingredientrepository

import (
	ingrediententity "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/entity"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
)

type IngredientRepository interface {
	Create(ingredient *ingrediententity.Ingredient) (*ingrediententity.Ingredient, errs.ErrMessage)
	Get() ([]ingrediententity.Ingredient, errs.ErrMessage)
	GetByName(name string) (*ingrediententity.Ingredient, errs.ErrMessage)
	GetById(id uuid.UUID) (*ingrediententity.Ingredient, errs.ErrMessage)
	Delete(id uuid.UUID, status bool) errs.ErrMessage
}
