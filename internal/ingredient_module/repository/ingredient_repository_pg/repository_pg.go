package ingredientrepositorypg

import (
	"fmt"

	ingrediententity "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/entity"
	ingredientrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/repository/ingredient_repository"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IngredientRepo struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) ingredientrepository.IngredientRepository {
	return &IngredientRepo{db: db}
}

func (r *IngredientRepo) Create(ingredient *ingrediententity.Ingredient) (*ingrediententity.Ingredient, errs.ErrMessage) {
	if err := r.db.Create(ingredient).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return ingredient, nil
}

func (r *IngredientRepo) Get() ([]ingrediententity.Ingredient, errs.ErrMessage) {
	var payload []ingrediententity.Ingredient

	if err := r.db.Where("is_active = true").
		Order("name ASC").
		Find(&payload).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	if len(payload) == 0 {
		return nil, errs.NewNotFound("No ingredients found")
	}

	return payload, nil
}

func (r *IngredientRepo) GetByName(name string) (*ingrediententity.Ingredient, errs.ErrMessage) {
	var payload ingrediententity.Ingredient

	if err := r.db.Where("name = ?", name).First(&payload).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFound(fmt.Sprintf("ingredient with name '%s' not found", name))
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &payload, nil
}

func (r *IngredientRepo) GetById(id uuid.UUID) (*ingrediententity.Ingredient, errs.ErrMessage) {
	var payload ingrediententity.Ingredient

	if err := r.db.Where("id = ? AND is_active = true", id).First(&payload).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFound("This ingredient not found or inactive")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &payload, nil
}

func (r *IngredientRepo) Delete(id uuid.UUID, status bool) errs.ErrMessage {
	res := r.db.Model(&ingrediententity.Ingredient{}).
		Where("id = ?", id).
		Update("is_active", status)

	if res.Error != nil {
		return errs.NewInternalServerError(res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return errs.NewNotFound("Ingredient not found")
	}

	return nil
}
