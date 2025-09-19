package categoryrepository

import (
	categoryentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/entity"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	Create(catgeory *categoryentity.Category) (*categoryentity.Category, errs.ErrMessage)
	Get() ([]categoryentity.Category, errs.ErrMessage)
	GetByName(name string) (*categoryentity.Category, errs.ErrMessage)
	GetById(id uuid.UUID) (*categoryentity.Category, errs.ErrMessage)
	Delete(id uuid.UUID, status bool) errs.ErrMessage
}
