package categoryrepository

import (
	categoryentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/entity"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	Create(catgeory *categoryentity.Category) (*categoryentity.Category, errs.ErrMessage)
	Get() ([]*categoryentity.Category, errs.ErrMessage)
	Update(id uuid.UUID, category *categoryentity.Category) (*categoryentity.Category, errs.ErrMessage)
}
