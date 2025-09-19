package categoryservice

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/dto"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
)

type CategoryService interface {
	Create(payload *dto.CreateCategoryRequest) (*dto.CategoryResponse, errs.ErrMessage)
	Get() ([]*dto.CategoryResponse, errs.ErrMessage)
	Delete(id uuid.UUID) errs.ErrMessage
}
