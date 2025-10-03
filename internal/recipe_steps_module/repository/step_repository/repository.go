package steprepository

import (
	stepdto "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/dto"
	stepentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/entity"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
)

type RecipeStepRepository interface {
	Create(step *stepentity.RecipeStep) (*stepentity.RecipeStep, errs.ErrMessage)
	Get() ([]stepentity.RecipeStep, errs.ErrMessage)
	GetById(RecipeId uuid.UUID) ([]stepentity.RecipeStep, errs.ErrMessage)
	Update(id uint, dto *stepdto.UpdateRequest) errs.ErrMessage
	Delete(id uint) errs.ErrMessage
	GetMaxStepNumberByRecipe(recipeId uuid.UUID) (int, errs.ErrMessage)
}
