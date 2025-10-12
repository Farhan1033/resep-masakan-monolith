package recipeserviceimp

import (
	"math"

	"github.com/Farhan1033/resep-masakan-monolith.git/event"
	authrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/repository/auth_repository"
	categoryrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/repository/category_repository"
	detailrecipeservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/service/detail_recipe_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/dto"
	recipeentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/entity"
	reciperepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/repository/recipe_repository"
	recipeservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/service/recipe_service"
	stepservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/service/step_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/logger"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type RecipeSvc struct {
	repoRecipe    reciperepository.RecipeRepository
	repoUser      authrepository.AuthRepostiory
	repoCategory  categoryrepository.CategoryRepository
	svcStep       stepservice.RecipeStepService
	svcIngredient detailrecipeservice.DetailRecipeService
	validate      *validator.Validate
	bus           *event.EventBus
}

func NewRecipeService(repoRecipe reciperepository.RecipeRepository,
	repoUser authrepository.AuthRepostiory,
	repoCategory categoryrepository.CategoryRepository,
	svcStep stepservice.RecipeStepService,
	svcIngredient detailrecipeservice.DetailRecipeService,
	bus *event.EventBus) recipeservice.RecipeService {
	return &RecipeSvc{
		repoRecipe:    repoRecipe,
		repoUser:      repoUser,
		repoCategory:  repoCategory,
		svcStep:       svcStep,
		svcIngredient: svcIngredient,
		validate:      validator.New(),
		bus:           bus,
	}
}

func (s *RecipeSvc) Create(payload *dto.CreateRequest, IdUser uuid.UUID) (*dto.CreateResponse, errs.ErrMessage) {
	if err := s.validate.Struct(payload); err != nil {
		formatError := validation.FormatValidationError(err)
		logger.Error("RecipeCreateValidationFailed", err, payload)
		return nil, errs.NewBadRequest(formatError.Message())
	}

	user, errUser := s.repoUser.GetById(IdUser)
	if errUser != nil {
		logger.Error("RecipeCreateUserNotFound", errUser, payload)
		return nil, errUser
	}

	category, errCategory := s.repoCategory.GetById(payload.CategoryId)
	if errCategory != nil {
		logger.Error("RecipeCreateCategoryNotFound", errCategory, payload)
		return nil, errCategory
	}

	request := &recipeentity.Recipe{
		CategoryId:     category.ID,
		Title:          payload.Title,
		Description:    payload.Description,
		DifficultLevel: payload.DifficultLevel,
		PrepTime:       payload.PrepTime,
		CookTime:       payload.CookTime,
		TotalTime:      payload.TotalTime,
		Servings:       payload.Servings,
		OriginRegion:   payload.OriginRegion,
		ImageUrl:       payload.ImageUrl,
		CreatedBy:      user.ID,
	}

	recipe, err := s.repoRecipe.Create(request)
	if err != nil {
		logger.Error("RecipeCreateFailed", err, payload)
		return nil, err
	}

	s.bus.Publish("recipeCreated", recipe)
	logger.Info("RecipeCreated", recipe)

	response := &dto.CreateResponse{
		ID:             recipe.ID,
		Title:          recipe.Title,
		Description:    recipe.Description,
		DifficultLevel: recipe.DifficultLevel,
		PrepTime:       recipe.PrepTime,
		CookTime:       recipe.CookTime,
		TotalTime:      recipe.TotalTime,
		Servings:       recipe.Servings,
		OriginRegion:   recipe.OriginRegion,
		ImageUrl:       recipe.ImageUrl,
		CreatedAt:      recipe.CreatedAt,
		Category: &dto.Category{
			ID:   category.ID,
			Name: category.Name,
		},
		CreatedBy: &dto.CreatedBy{
			ID:       user.ID,
			FullName: user.FullName,
		},
	}
	
	return response, nil
}

func (s *RecipeSvc) GetByPagination(page, limit int) (*dto.PaginatedResponse, errs.ErrMessage) {
	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	recipes, total, err := s.repoRecipe.GetByPagination(limit, offset)
	if err != nil {
		return nil, err
	}

	if len(recipes) == 0 {
		return nil, errs.NewNotFound("Recipe not found!")
	}

	response := make([]*dto.GetResponse, len(recipes))
	for i, recipe := range recipes {
		response[i] = &dto.GetResponse{
			ID:             recipe.ID,
			Title:          recipe.Title,
			Description:    recipe.Description,
			DifficultLevel: recipe.DifficultLevel,
			PrepTime:       recipe.PrepTime,
			CookTime:       recipe.CookTime,
			TotalTime:      recipe.CookTime,
			Servings:       recipe.Servings,
			OriginRegion:   recipe.OriginRegion,
			ImageUrl:       recipe.ImageUrl,
			Category: &dto.Category{
				ID:   recipe.CategoryID,
				Name: recipe.CategoryName,
			},
			CreatedAt: recipe.CreatedAt,
			UpdateAt:  recipe.UpdatedAt,
			CreatedBy: &dto.CreatedBy{
				ID:       recipe.UserID,
				FullName: recipe.UserName,
			},
		}
	}

	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	return &dto.PaginatedResponse{
		Data:        response,
		CurrentPage: page,
		PerPage:     limit,
		Total:       total,
		TotalPages:  totalPage,
		HasNext:     page < totalPage,
		HasPrev:     page > 1,
	}, nil
}

func (s *RecipeSvc) GetById(id uuid.UUID) (*dto.GetResponse, errs.ErrMessage) {
	recipe, err := s.repoRecipe.GetById(id)
	if err != nil {
		return nil, err
	}

	return &dto.GetResponse{
		ID:             recipe.ID,
		Title:          recipe.Title,
		Description:    recipe.Description,
		DifficultLevel: recipe.DifficultLevel,
		PrepTime:       recipe.PrepTime,
		CookTime:       recipe.CookTime,
		TotalTime:      recipe.TotalTime,
		Servings:       recipe.Servings,
		OriginRegion:   recipe.OriginRegion,
		ImageUrl:       recipe.ImageUrl,
		Category: &dto.Category{
			ID:   recipe.CategoryID,
			Name: recipe.CategoryName,
		},
		CreatedAt: recipe.CreatedAt,
		UpdateAt:  recipe.UpdatedAt,
		CreatedBy: &dto.CreatedBy{
			ID:       recipe.UserID,
			FullName: recipe.UserName,
		},
	}, nil
}

func (s *RecipeSvc) Update(id uuid.UUID, payload *dto.UpdateRequest, userId uuid.UUID) (*dto.UpdateResponse, errs.ErrMessage) {
	existingRecipe, err := s.repoRecipe.GetById(id)
	if err != nil {
		logger.Error("RecipeUpdateNotFound", err, id)
		return nil, err
	}

	if existingRecipe.UserID != userId {
		errForbidden := errs.NewForbidden("You are not allowed to update this recipe")
		logger.Error("RecipeUpdateForbidden", errForbidden, id)
		return nil, errForbidden
	}

	updateData := &recipeentity.Recipe{
		CategoryId:     payload.CategoryId,
		Title:          payload.Title,
		Description:    payload.Description,
		DifficultLevel: payload.DifficultLevel,
		PrepTime:       payload.PrepTime,
		CookTime:       payload.CookTime,
		TotalTime:      payload.TotalTime,
		Servings:       payload.Servings,
		OriginRegion:   payload.OriginRegion,
		ImageUrl:       payload.ImageUrl,
	}

	updated, errUpdate := s.repoRecipe.Update(id, updateData)
	if errUpdate != nil {
		logger.Error("RecipeUpdateFailed", errUpdate, payload)
		return nil, errUpdate
	}

	s.bus.Publish("recipeUpdated", updated)
	logger.Info("RecipeUpdated", updated)

	return &dto.UpdateResponse{
		ID:             updated.ID,
		CategoryId:     updated.CategoryId,
		Title:          updated.Title,
		Description:    updated.Description,
		DifficultLevel: updated.DifficultLevel,
		PrepTime:       updated.PrepTime,
		CookTime:       updated.CookTime,
		TotalTime:      updated.TotalTime,
		Servings:       updated.Servings,
		OriginRegion:   updated.OriginRegion,
		ImageUrl:       updated.ImageUrl,
		UpdatedAt:      updated.UpdatedAt,
	}, nil
}

func (s *RecipeSvc) Delete(id uuid.UUID, status bool) errs.ErrMessage {
	err := s.repoRecipe.Delete(id, status)
	if err != nil {
		logger.Error("RecipeDeleteFailed", err, id)
		return err
	}

	s.bus.Publish("recipeDeleted", id)
	logger.Info("RecipeDeleted", id)

	return nil
}

func (s *RecipeSvc) GetDetailRecipe(id uuid.UUID) (*dto.RecipeDetail, errs.ErrMessage) {
	recipe, errMsgRecipe := s.repoRecipe.GetById(id)
	if errMsgRecipe != nil {
		return nil, errMsgRecipe
	}

	steps, errMsgStep := s.svcStep.GetByRecipeId(id)
	if errMsgStep != nil {
		return nil, errMsgStep
	}

	ingredients, errMsgIngredient := s.svcIngredient.GetDetailRecipeByRecipeId(id)
	if errMsgIngredient != nil {
		return nil, errMsgIngredient
	}

	response := &dto.RecipeDetail{
		ID:             recipe.ID,
		Title:          recipe.Title,
		Description:    recipe.Description,
		DifficultLevel: recipe.DifficultLevel,
		PrepTime:       recipe.PrepTime,
		CookTime:       recipe.CookTime,
		TotalTime:      recipe.TotalTime,
		Servings:       recipe.Servings,
		OriginRegion:   recipe.OriginRegion,
		ImageUrl:       recipe.ImageUrl,
		IsActive:       recipe.IsActive,
		Categories: &dto.Category{
			ID:   recipe.CategoryID,
			Name: recipe.CategoryName,
		},
		Steps:       steps,
		Ingredients: ingredients,
		CreatedBy: &dto.CreatedBy{
			ID:       recipe.UserID,
			FullName: recipe.UserName,
		},
		CreatedAt: recipe.CreatedAt,
		UpdatedAt: recipe.UpdatedAt,
	}

	return response, nil
}
