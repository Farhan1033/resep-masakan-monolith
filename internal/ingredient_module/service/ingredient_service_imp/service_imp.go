package ingredientserviceimp

import (
	"fmt"

	"github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/dto"
	ingrediententity "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/entity"
	ingredientrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/repository/ingredient_repository"
	ingredientservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/service/ingredient_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IngredientSvc struct {
	repo     ingredientrepository.IngredientRepository
	validate *validator.Validate
}

func NewIngredientService(repo ingredientrepository.IngredientRepository) ingredientservice.IngredientService {
	return &IngredientSvc{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *IngredientSvc) Create(payload *dto.CreateRequest) (*dto.CreateResponse, errs.ErrMessage) {
	if err := s.validate.Struct(payload); err != nil {
		formatError := validation.FormatValidationError(err)
		return nil, errs.NewBadRequest(fmt.Sprintf("Required: %s", formatError))
	}

	if _, err := s.repo.GetByName(payload.Name); err == nil {
		return nil, errs.NewFound(fmt.Sprintf("This ingredient %s already exist!", payload.Name))
	}

	ingredientNew := &ingrediententity.Ingredient{
		Name: payload.Name,
	}

	CreateIngredient, err := s.repo.Create(ingredientNew)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	response := &dto.CreateResponse{
		ID:        CreateIngredient.ID,
		Name:      CreateIngredient.Name,
		IsActive:  CreateIngredient.IsActive,
		CreatedAt: CreateIngredient.CreatedAt,
		UpdatedAt: CreateIngredient.UpdatedAt,
	}

	return response, nil
}

func (s *IngredientSvc) Get() ([]*dto.CreateResponse, errs.ErrMessage) {
	ingredients, err := s.repo.Get()
	if err != nil {
		return nil, errs.NewNotFound(err.Error())
	}

	response := make([]*dto.CreateResponse, len(ingredients))
	for i, ingredient := range ingredients {
		response[i] = &dto.CreateResponse{
			ID:        ingredient.ID,
			Name:      ingredient.Name,
			IsActive:  ingredient.IsActive,
			CreatedAt: ingredient.CreatedAt,
			UpdatedAt: ingredient.UpdatedAt,
		}
	}

	return response, nil
}

func (s *IngredientSvc) Delete(id uuid.UUID) errs.ErrMessage {
	response := s.repo.Delete(id, false)
	if response != nil {
		return errs.NewInternalServerError(response.Error())
	}

	return nil
}
