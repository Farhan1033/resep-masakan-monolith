package recipeservice

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/dto"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
)

type RecipeService interface {
	Create(payload *dto.CreateRequest, IdUser uuid.UUID) (*dto.CreateResponse, errs.ErrMessage)
	GetByPagination(page, limit int) (*dto.PaginatedResponse, errs.ErrMessage)
	GetById(id uuid.UUID) (*dto.GetResponse, errs.ErrMessage)
	Update(id uuid.UUID, payload *dto.UpdateRequest, userId uuid.UUID) (*dto.UpdateResponse, errs.ErrMessage)
	Delete(id uuid.UUID, status bool) errs.ErrMessage
	GetDetailRecipe(id uuid.UUID) (*dto.RecipeDetail, errs.ErrMessage)
}
