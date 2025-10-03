package stepserviceimpl

import (
	stepdto "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/dto"
	stepentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/entity"
	steprepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/repository/step_repository"
	stepservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/service/step_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type recipeStepSvc struct {
	repo     steprepository.RecipeStepRepository
	validate *validator.Validate
}

func NewRecipeStepService(repo steprepository.RecipeStepRepository) stepservice.RecipeStepService {
	return &recipeStepSvc{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *recipeStepSvc) Create(payload *stepdto.CreateRequest) (*stepdto.CreateResponse, errs.ErrMessage) {
	if err := s.validate.Struct(payload); err != nil {
		formatError := validation.FormatValidationError(err)
		return nil, errs.NewBadRequest(formatError.Message())
	}

	maxStep, errMsg := s.repo.GetMaxStepNumberByRecipe(payload.RecipeId)
	if errMsg != nil {
		return nil, errMsg
	}

	newStep := &stepentity.RecipeStep{
		RecipeId:    payload.RecipeId,
		Instruction: payload.Instruction,
		StepNumber:  maxStep + 1,
	}

	created, errMsg := s.repo.Create(newStep)
	if errMsg != nil {
		return nil, errMsg
	}

	res := &stepdto.CreateResponse{
		ID:          created.ID,
		RecipeId:    created.RecipeId,
		StepNumber:  created.StepNumber,
		Instruction: created.Instruction,
		CreatedAt:   created.CreatedAt,
	}

	return res, nil
}

func (s *recipeStepSvc) Get() ([]stepdto.RecipeStepResponse, errs.ErrMessage) {
	steps, err := s.repo.Get()
	if err != nil {
		return nil, err
	}

	var result []stepdto.RecipeStepResponse
	for _, step := range steps {
		result = append(result, stepdto.RecipeStepResponse{
			ID:          step.ID,
			RecipeId:    step.RecipeId,
			StepNumber:  step.StepNumber,
			Instruction: step.Instruction,
			CreatedAt:   step.CreatedAt,
			UpdatedAt:   step.UpdatedAt,
		})
	}

	return result, nil
}

func (s *recipeStepSvc) GetByRecipeId(recipeId uuid.UUID) ([]*stepdto.RecipeStepResponse, errs.ErrMessage) {
	if recipeId == uuid.Nil {
		return nil, errs.NewBadRequest("Invalid recipe ID")
	}

	steps, err := s.repo.GetById(recipeId)
	if err != nil {
		return nil, err
	}

	result := make([]*stepdto.RecipeStepResponse, 0, len(steps))
	for _, step := range steps {
		result = append(result, &stepdto.RecipeStepResponse{
			ID:          step.ID,
			RecipeId:    step.RecipeId,
			StepNumber:  step.StepNumber,
			Instruction: step.Instruction,
			CreatedAt:   step.CreatedAt,
			UpdatedAt:   step.UpdatedAt,
		})
	}

	return result, nil
}

func (s *recipeStepSvc) Update(id uint, payload *stepdto.UpdateRequest) errs.ErrMessage {
	if err := s.validate.Struct(payload); err != nil {
		formatError := validation.FormatValidationError(err)
		return errs.NewBadRequest(formatError.Message())
	}

	err := s.repo.Update(id, payload)
	if err != nil {
		return err
	}

	return err
}

func (s *recipeStepSvc) Delete(id uint) errs.ErrMessage {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return err
}
