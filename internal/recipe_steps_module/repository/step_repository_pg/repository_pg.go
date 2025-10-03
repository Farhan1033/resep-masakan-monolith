package steprepositorypg

import (
	stepdto "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/dto"
	stepentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/entity"
	steprepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/repository/step_repository"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type recipeStepRepo struct {
	db *gorm.DB
}

func NewRecipeStepRepository(db *gorm.DB) steprepository.RecipeStepRepository {
	return &recipeStepRepo{
		db: db,
	}
}

func (r *recipeStepRepo) Create(step *stepentity.RecipeStep) (*stepentity.RecipeStep, errs.ErrMessage) {
	if err := r.db.Create(step).Error; err != nil {
		return nil, errs.NewInternalServerError("Failed to create recipe step: " + err.Error())
	}
	return step, nil
}

func (r *recipeStepRepo) Get() ([]stepentity.RecipeStep, errs.ErrMessage) {
	var steps []stepentity.RecipeStep
	if err := r.db.Find(&steps).Error; err != nil {
		return nil, errs.NewInternalServerError("Failed to fetch recipe steps: " + err.Error())
	}
	return steps, nil
}

func (r *recipeStepRepo) Update(id uint, dto *stepdto.UpdateRequest) errs.ErrMessage {
	if err := r.db.Model(&stepentity.RecipeStep{}).
		Where("id = ?", id).
		Updates(dto).Error; err != nil {
		return errs.NewInternalServerError("Failed to update recipe step: " + err.Error())
	}

	return nil
}

func (r *recipeStepRepo) Delete(id uint) errs.ErrMessage {
	if err := r.db.Delete(&stepentity.RecipeStep{}, id).Error; err != nil {
		return errs.NewInternalServerError("Failed to delete recipe step: " + err.Error())
	}
	return nil
}

func (r *recipeStepRepo) GetMaxStepNumberByRecipe(recipeId uuid.UUID) (int, errs.ErrMessage) {
	var maxStep int
	err := r.db.Model(&stepentity.RecipeStep{}).
		Where("recipe_id = ?", recipeId).
		Select("COALESCE(MAX(step_number), 0)").Scan(&maxStep).Error
	if err != nil {
		return 0, errs.NewInternalServerError("failed to get max step number: " + err.Error())
	}
	return maxStep, nil
}

func (r *recipeStepRepo) GetById(RecipeId uuid.UUID) ([]stepentity.RecipeStep, errs.ErrMessage) {
	var steps []stepentity.RecipeStep

	if err := r.db.Where("recipe_id = ?", RecipeId).Order("step_number ASC").Find(&steps).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	if len(steps) == 0 {
		return nil, errs.NewNotFound("Steps not found for this recipe")
	}

	return steps, nil
}
