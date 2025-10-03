package detailreciperepositorypg

import (
	"fmt"

	"github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/dto"
	detailrecipeentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/entity"
	detailreciperepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/repository/detail_recipe_repository"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DetailRecipeRepo struct {
	db *gorm.DB
}

func NewDetailRecipeRepository(db *gorm.DB) detailreciperepository.DetailRecipeRepository {
	return &DetailRecipeRepo{
		db: db,
	}
}

func (r *DetailRecipeRepo) Create(detail *detailrecipeentity.DetailRecipeEntity) (*detailrecipeentity.DetailRecipeEntity, errs.ErrMessage) {
	if err := r.db.Create(detail).Error; err != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to create detail recipe: %v", err))
	}
	return detail, nil
}

func (r *DetailRecipeRepo) Get() ([]dto.RecipeWithIngredients, errs.ErrMessage) {
	var results []dto.RecipeWithIngredients

	if err := r.db.Table("detail_recipes AS ri").
		Select(`
            ri.id,
            ri.recipe_id AS recipe_id,
            i.id AS ingredient_id,
            i.name AS ingredient_name,
            ri.quantity,
            ri.unit,
			ri.created_by
        `).
		Joins("JOIN ingredient i ON i.id = ri.ingredient_id").
		Scan(&results).Error; err != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to get detail recipes: %v", err))
	}

	// kalau kosong, tetap return slice kosong (bukan error)
	return results, nil
}

func (r *DetailRecipeRepo) GetById(id uuid.UUID) (*detailrecipeentity.DetailRecipeEntity, errs.ErrMessage) {
	var result detailrecipeentity.DetailRecipeEntity

	if err := r.db.First(&result, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFound(fmt.Sprintf("detail recipe with id '%s' not found", id))
		}
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to get detail recipe by id: %v", err))
	}

	return &result, nil
}

func (r *DetailRecipeRepo) GetByRecipeId(idRecipe uuid.UUID) ([]dto.RecipeWithIngredients, errs.ErrMessage) {
	var result []dto.RecipeWithIngredients

	if err := r.db.Table("detail_recipes AS ri").
		Select(`
            ri.id,
            ri.recipe_id AS recipe_id,
            i.id AS ingredient_id,
            i.name AS ingredient_name,
            ri.quantity,
            ri.unit,
			ri.created_by
        `).
		Joins("JOIN ingredient i ON i.id = ri.ingredient_id").
		Where("ri.recipe_id = ?", idRecipe).
		Scan(&result).Error; err != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to get detail recipes by recipe id: %v", err))
	}

	return result, nil
}

func (r *DetailRecipeRepo) Update(id uuid.UUID, detail *detailrecipeentity.DetailRecipeEntity) errs.ErrMessage {
	update := map[string]interface{}{
		"recipe_id":     detail.RecipeID,
		"ingredient_id": detail.IngredientID,
		"quantity":      detail.Quantity,
		"unit":          detail.Unit,
	}

	result := r.db.Model(&detailrecipeentity.DetailRecipeEntity{}).Where("id = ?", id).Updates(update)
	if result.Error != nil {
		return errs.NewInternalServerError(fmt.Sprintf("failed to update detail recipe: %v", result.Error))
	}
	if result.RowsAffected == 0 {
		return errs.NewNotFound(fmt.Sprintf("detail recipe with id '%s' not found", id))
	}

	return nil
}

func (r *DetailRecipeRepo) Delete(id uuid.UUID) errs.ErrMessage {
	result := r.db.Delete(&detailrecipeentity.DetailRecipeEntity{}, "id = ?", id)
	if result.Error != nil {
		return errs.NewInternalServerError(fmt.Sprintf("failed to delete detail recipe: %v", result.Error))
	}
	if result.RowsAffected == 0 {
		return errs.NewNotFound(fmt.Sprintf("detail recipe with id '%s' not found", id))
	}
	return nil
}
