package reciperepository

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/dto"
	recipeentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/entity"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
)

type RecipeRepository interface {
	Create(recipe *recipeentity.Recipe) (*recipeentity.Recipe, errs.ErrMessage)
	GetByPagination(limit, offset int) ([]dto.RecipeWithRelations, int, errs.ErrMessage)
	GetById(id uuid.UUID) (*dto.RecipeWithRelations, errs.ErrMessage)
	Update(id uuid.UUID, recipe *recipeentity.Recipe) (*recipeentity.Recipe, errs.ErrMessage)
	Delete(id uuid.UUID, status bool) errs.ErrMessage
}
