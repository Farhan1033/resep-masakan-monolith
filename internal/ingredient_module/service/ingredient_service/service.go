package ingredientservice

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/dto"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
)

type IngredientService interface {
	Create(payload *dto.CreateRequest) (*dto.CreateResponse, errs.ErrMessage)
	Get() ([]*dto.CreateResponse, errs.ErrMessage)
	Delete(id uuid.UUID) errs.ErrMessage
}
