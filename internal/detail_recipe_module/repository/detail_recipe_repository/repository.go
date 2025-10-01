package detailreciperepository

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/dto"
	detailrecipeentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/entity"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
)

type DetailRecipeRepository interface {
	Get() ([]dto.RecipeWithIngredients, errs.ErrMessage)
	GetById(id uuid.UUID) (*dto.RecipeWithIngredients, errs.ErrMessage)
	GetByRecipeId(idRecipe uuid.UUID) (*dto.RecipeWithIngredients, errs.ErrMessage)
	Create(detail *detailrecipeentity.DetailRecipeEntity) (*detailrecipeentity.DetailRecipeEntity, errs.ErrMessage)
	Update(id uuid.UUID, detail *detailrecipeentity.DetailRecipeEntity) errs.ErrMessage
	Delete(id uuid.UUID) errs.ErrMessage
}
