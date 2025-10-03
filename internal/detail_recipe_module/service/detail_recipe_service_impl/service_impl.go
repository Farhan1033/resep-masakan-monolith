package detailrecipeserviceimpl

import (
	"time"

	authrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/repository/auth_repository"
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/dto"
	detailrecipeentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/entity"
	detailreciperepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/repository/detail_recipe_repository"
	detailrecipeservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/service/detail_recipe_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type DetailRecipeSvc struct {
	repo     detailreciperepository.DetailRecipeRepository
	repoUser authrepository.AuthRepostiory
	validate *validator.Validate
}

func NewDetailRecipeService(repo detailreciperepository.DetailRecipeRepository,
	repoUser authrepository.AuthRepostiory) detailrecipeservice.DetailRecipeService {
	return &DetailRecipeSvc{
		repo:     repo,
		repoUser: repoUser,
		validate: validator.New(),
	}
}

func (s *DetailRecipeSvc) CreateDetailRecipe(id uuid.UUID, req *dto.CreateDetailRecipeRequest) (*dto.DetailRecipeResponse, errs.ErrMessage) {
	if err := s.validate.Struct(req); err != nil {
		formatError := validation.FormatValidationError(err)
		return nil, errs.NewBadRequest(formatError.Message())
	}

	getUser, errUser := s.repoUser.GetById(id)
	if errUser != nil {
		return nil, errs.NewNotFound(errUser.Message())
	}

	entity := &detailrecipeentity.DetailRecipeEntity{
		RecipeID:     req.RecipeID,
		IngredientID: req.IngredientID,
		Quantity:     req.Quantity,
		Unit:         req.Unit,
		CreatedBy:    getUser.FullName,
	}

	result, err := s.repo.Create(entity)
	if err != nil {
		return nil, err
	}

	response := &dto.DetailRecipeResponse{
		ID:           result.ID,
		RecipeID:     result.RecipeID,
		IngredientID: result.IngredientID,
		Quantity:     result.Quantity,
		Unit:         result.Unit,
		CreatedAt:    time.Now(),
		CreatedBy:    getUser.FullName,
	}

	return response, nil
}

func (s *DetailRecipeSvc) GetAllDetailRecipes() ([]dto.RecipeWithIngredientsResponse, errs.ErrMessage) {
	results, err := s.repo.Get()
	if err != nil {
		return nil, errs.NewInternalServerError(err.Message())
	}

	recipeMap := make(map[uuid.UUID]*dto.RecipeWithIngredientsResponse)

	for _, item := range results {
		if _, exists := recipeMap[item.RecipeID]; !exists {
			recipeMap[item.RecipeID] = &dto.RecipeWithIngredientsResponse{
				RecipeID:    item.RecipeID,
				Ingredients: []dto.IngredientDetail{},
			}
		}

		ingredient := dto.IngredientDetail{
			ID:             item.ID,
			IngredientID:   item.IngredientID,
			IngredientName: item.IngredientName,
			Quantity:       item.Quantity,
			Unit:           item.Unit,
		}

		recipeMap[item.RecipeID].Ingredients = append(recipeMap[item.RecipeID].Ingredients, ingredient)
	}

	var response []dto.RecipeWithIngredientsResponse
	for _, recipe := range recipeMap {
		response = append(response, *recipe)
	}

	return response, nil
}

func (s *DetailRecipeSvc) GetDetailRecipeById(id uuid.UUID) (*dto.DetailRecipeResponse, errs.ErrMessage) {
	result, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Message())
	}

	response := &dto.DetailRecipeResponse{
		ID:           result.ID,
		RecipeID:     result.RecipeID,
		IngredientID: result.IngredientID,
		Quantity:     result.Quantity,
		Unit:         result.Unit,
		CreatedAt:    result.CreatedAt,
		CreatedBy:    result.CreatedBy,
	}

	return response, nil
}

func (s *DetailRecipeSvc) GetDetailRecipeByRecipeId(recipeId uuid.UUID) (*dto.RecipeWithIngredientsResponse, errs.ErrMessage) {
	if recipeId == uuid.Nil {
		return nil, errs.NewBadRequest("Invalid Recipe ID")
	}

	result, err := s.repo.GetByRecipeId(recipeId)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errs.NewNotFound("Recipe not found")
	}

	firstItem := result[0]

	ingredients := make([]dto.IngredientDetail, 0, len(result))
	for _, item := range result {
		ingredient := dto.IngredientDetail{
			ID:             firstItem.ID,
			IngredientID:   item.IngredientID,
			IngredientName: item.IngredientName,
			Quantity:       item.Quantity,
			Unit:           item.Unit,
		}

		ingredients = append(ingredients, ingredient)
	}

	response := &dto.RecipeWithIngredientsResponse{
		RecipeID:    firstItem.RecipeID,
		Ingredients: ingredients,
	}

	return response, nil
}

func (s *DetailRecipeSvc) UpdateDetailRecipe(id uuid.UUID, req *dto.UpdateDetailRecipeRequest) errs.ErrMessage {
	if err := s.validate.Struct(req); err != nil {
		formatError := validation.FormatValidationError(err)
		return errs.NewBadRequest(formatError.Message())
	}

	_, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	entity := &detailrecipeentity.DetailRecipeEntity{
		RecipeID:     req.RecipeID,
		IngredientID: req.IngredientID,
		Quantity:     req.Quantity,
		Unit:         req.Unit,
	}

	return s.repo.Update(id, entity)
}

func (s *DetailRecipeSvc) DeleteDetailRecipe(id uuid.UUID) errs.ErrMessage {
	_, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
