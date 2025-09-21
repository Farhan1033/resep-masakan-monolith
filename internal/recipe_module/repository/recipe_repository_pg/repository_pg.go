package reciperepositorypg

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/dto"
	recipeentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/entity"
	reciperepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/repository/recipe_repository"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RecipeRepo struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) reciperepository.RecipeRepository {
	return &RecipeRepo{
		db: db,
	}
}

func (r *RecipeRepo) Create(recipe *recipeentity.Recipe) (*recipeentity.Recipe, errs.ErrMessage) {
	if err := r.db.Create(recipe).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return recipe, nil
}

func (r *RecipeRepo) GetByPagination(limit, offset int) ([]dto.RecipeWithRelations, int, errs.ErrMessage) {
	var recipes []dto.RecipeWithRelations
	var total int64

	if err := r.db.Table("recipe").Count(&total).Error; err != nil {
		return nil, 0, errs.NewNotFound("Data not found!")
	}

	if err := r.db.Table("recipe").
		Select(`
            recipe.id,
            recipe.title,
            recipe.description,
            recipe.difficult_level,
            recipe.prep_time,
            recipe.cook_time,
            recipe.total_time,
            recipe.servings,
            recipe.origin_region,
            recipe.image_url,
            recipe.is_active,
			recipe.created_at,
			recipe.updated_at,
            users.id AS user_id,
            users.full_name AS user_name,
            category.id AS category_id,
            category.name AS category_name
        `).
		Joins("LEFT JOIN users ON users.id = recipe.created_by").
		Joins("LEFT JOIN category ON category.id = recipe.category_id").
		Where("recipe.is_active = ?", true).
		Limit(limit).
		Offset(offset).
		Scan(&recipes).Error; err != nil {
		return nil, 0, errs.NewInternalServerError(err.Error())
	}

	return recipes, int(total), nil
}

func (r *RecipeRepo) GetById(id uuid.UUID) (*dto.RecipeWithRelations, errs.ErrMessage) {
	var recipe dto.RecipeWithRelations

	if err := r.db.Table("recipe").
		Select(`
            recipe.id,
            recipe.title,
            recipe.description,
            recipe.difficult_level,
            recipe.prep_time,
            recipe.cook_time,
            recipe.total_time,
            recipe.servings,
            recipe.origin_region,
            recipe.image_url,
            recipe.is_active,
			recipe.created_at,
			recipe.updated_at,
            users.id AS user_id,
            users.full_name AS user_name,
            category.id AS category_id,
            category.name AS category_name
        `).
		Joins("LEFT JOIN users ON users.id = recipe.created_by").
		Joins("LEFT JOIN category ON category.id = recipe.category_id").
		Where("recipe.id = ?", id).
		Scan(&recipe).Error; err != nil {
		return nil, errs.NewNotFound("Data not found!")
	}

	return &recipe, nil
}

func (r *RecipeRepo) Update(id uuid.UUID, recipe *recipeentity.Recipe) (*recipeentity.Recipe, errs.ErrMessage) {
	updates := map[string]interface{}{
		"category_id":     recipe.CategoryId,
		"title":           recipe.Title,
		"description":     recipe.Description,
		"difficult_level": recipe.DifficultLevel,
		"prep_time":       recipe.PrepTime,
		"cook_time":       recipe.CookTime,
		"total_time":      recipe.TotalTime,
		"servings":        recipe.Servings,
		"origin_region":   recipe.OriginRegion,
		"image_url":       recipe.ImageUrl,
	}

	result := r.db.Model(&recipeentity.Recipe{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return nil, errs.NewInternalServerError(result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, errs.NewNotFound("Recipe not found")
	}

	var updated recipeentity.Recipe
	if err := r.db.Where("id = ?", id).First(&updated).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &updated, nil
}

func (r *RecipeRepo) Delete(id uuid.UUID, status bool) errs.ErrMessage {
	if err := r.db.Model(&recipeentity.Recipe{}).
		Where("id = ?", id).
		Update("is_active", status).Error; err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}
