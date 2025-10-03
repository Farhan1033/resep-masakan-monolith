package stepservice

import (
	stepdto "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/dto"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
)

type RecipeStepService interface {
	Create(payload *stepdto.CreateRequest) (*stepdto.CreateResponse, errs.ErrMessage)
	Get() ([]stepdto.RecipeStepResponse, errs.ErrMessage)
	Update(id uint, payload *stepdto.UpdateRequest) errs.ErrMessage
	Delete(id uint) errs.ErrMessage
}
