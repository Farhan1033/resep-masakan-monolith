package detailrecipeservice

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/dto"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
)

type DetailRecipeService interface {
	CreateDetailRecipe(id uuid.UUID, req *dto.CreateDetailRecipeRequest) (*dto.DetailRecipeResponse, errs.ErrMessage)
	GetAllDetailRecipes() ([]dto.RecipeWithIngredientsResponse, errs.ErrMessage)
	GetDetailRecipeById(id uuid.UUID) (*dto.DetailRecipeResponse, errs.ErrMessage)
	GetDetailRecipeByRecipeId(recipeId uuid.UUID) (*dto.RecipeWithIngredientsResponse, errs.ErrMessage)
	UpdateDetailRecipe(id uuid.UUID, req *dto.UpdateDetailRecipeRequest) errs.ErrMessage
	DeleteDetailRecipe(id uuid.UUID) errs.ErrMessage
}
