package detailreciperepositorypg

import (
	"errors"

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
		return nil, errs.NewInternalServerError(err.Error())
	}
	return detail, nil
}

func (r *DetailRecipeRepo) Get() ([]dto.RecipeWithIngredients, errs.ErrMessage) {
	var results []dto.RecipeWithIngredients

	if err := r.db.Table("detail_recipes AS ri").
		Select(`
            ri.id,
            r.id AS recipe_id,
            r.title AS recipe_title,
            r.description AS recipe_description,
            r.difficult_level AS difficult_level,
            r.prep_time AS prep_time,
            r.cook_time AS cook_time,
            r.total_time AS total_time,
            r.servings AS servings,
            r.origin_region AS origin_region,
            r.image_url AS image_url,
            r.is_active AS is_active,
            i.id AS ingredient_id,
            i.name AS ingredient_name,
            ri.quantity,
            ri.unit,
			ri.created_by
        `).
		Joins("JOIN recipe r ON r.id = ri.recipe_id").
		Joins("JOIN ingredient i ON i.id = ri.ingredient_id").
		Scan(&results).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFound("Data not found!")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	return results, nil
}

func (r *DetailRecipeRepo) GetById(id uuid.UUID) (*dto.RecipeWithIngredients, errs.ErrMessage) {
	var result dto.RecipeWithIngredients

	if err := r.db.Table("detail_recipes AS ri").
		Select(`
            ri.id,
            r.id AS recipe_id,
            r.title AS recipe_title,
            r.description AS recipe_description,
            r.difficult_level AS difficult_level,
            r.prep_time AS prep_time,
            r.cook_time AS cook_time,
            r.total_time AS total_time,
            r.servings AS servings,
            r.origin_region AS origin_region,
            r.image_url AS image_url,
            r.is_active AS is_active,
            i.id AS ingredient_id,
            i.name AS ingredient_name,
            ri.quantity,
            ri.unit,
			ri.created_by
        `).
		Joins("JOIN recipe r ON r.id = ri.recipe_id").
		Joins("JOIN ingredient i ON i.id = ri.ingredient_id").
		Where("ri.id = ?", id).
		Scan(&result).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFound("Detail recipe not found!")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *DetailRecipeRepo) GetByRecipeId(idRecipe uuid.UUID) (*dto.RecipeWithIngredients, errs.ErrMessage) {
	var result dto.RecipeWithIngredients

	if err := r.db.Table("detail_recipes AS ri").
		Select(`
            ri.id,
            r.id AS recipe_id,
            r.title AS recipe_title,
            r.description AS recipe_description,
            r.difficult_level AS difficult_level,
            r.prep_time AS prep_time,
            r.cook_time AS cook_time,
            r.total_time AS total_time,
            r.servings AS servings,
            r.origin_region AS origin_region,
            r.image_url AS image_url,
            r.is_active AS is_active,
            i.id AS ingredient_id,
            i.name AS ingredient_name,
            ri.quantity,
            ri.unit,
			ri.created_by
        `).
		Joins("JOIN recipe r ON r.id = ri.recipe_id").
		Joins("JOIN ingredient i ON i.id = ri.ingredient_id").
		Where("r.id = ?", idRecipe).
		Scan(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFound("Recipe not found!")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *DetailRecipeRepo) Update(id uuid.UUID, detail *detailrecipeentity.DetailRecipeEntity) errs.ErrMessage {
	update := map[string]interface{}{
		"recipe_id":     detail.RecipeID,
		"ingredient_id": detail.IngredientID,
		"quantity":      detail.Quantity,
		"unit":          detail.Unit,
	}

	if err := r.db.Model(&detailrecipeentity.DetailRecipeEntity{}).Where("id = ?", id).
		Updates(update).Error; err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *DetailRecipeRepo) Delete(id uuid.UUID) errs.ErrMessage {
	if err := r.db.Delete(&detailrecipeentity.DetailRecipeEntity{}, "id = ?", id).Error; err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	return nil
}
