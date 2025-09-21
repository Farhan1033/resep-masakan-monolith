package recipeserviceimp

import (
	"math"

	authrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/repository/auth_repository"
	categoryrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/repository/category_repository"
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/dto"
	recipeentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/entity"
	reciperepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/repository/recipe_repository"
	recipeservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/service/recipe_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type RecipeSvc struct {
	repoRecipe   reciperepository.RecipeRepository
	repoUser     authrepository.AuthRepostiory
	repoCategory categoryrepository.CategoryRepository
	validate     *validator.Validate
}

func NewRecipeService(repoRecipe reciperepository.RecipeRepository,
	repoUser authrepository.AuthRepostiory,
	repoCategory categoryrepository.CategoryRepository) recipeservice.RecipeService {
	return &RecipeSvc{
		repoRecipe:   repoRecipe,
		repoUser:     repoUser,
		repoCategory: repoCategory,
		validate:     validator.New(),
	}
}

func (s *RecipeSvc) Create(payload *dto.CreateRequest, IdUser uuid.UUID) (*dto.CreateResponse, errs.ErrMessage) {
	if err := s.validate.Struct(payload); err != nil {
		formatError := validation.FormatValidationError(err)
		return nil, errs.NewBadRequest(formatError.Message())
	}

	user, errUser := s.repoUser.GetById(IdUser)
	if errUser != nil {
		return nil, errs.NewNotFound(errUser.Error())
	}

	userId := &dto.CreatedBy{
		ID:       user.ID,
		FullName: user.FullName,
	}

	category, errCategory := s.repoCategory.GetById(payload.CategoryId)

	if errCategory != nil {
		return nil, errs.NewNotFound(errCategory.Error())
	}

	categoryId := &dto.Category{
		ID:   category.ID,
		Name: category.Name,
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
		return nil, errs.NewInternalServerError(err.Error())
	}

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
		Category:       categoryId,
		CreatedBy:      userId,
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
		return nil, errs.NewInternalServerError(err.Error())
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
		return nil, errs.NewNotFound("Recipe not found!")
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
	if err := s.validate.Struct(payload); err != nil {
		formatError := validation.FormatValidationError(err)
		return nil, errs.NewBadRequest(formatError.Message())
	}

	existingRecipe, err := s.repoRecipe.GetById(id)
	if err != nil {
		return nil, err
	}

	if existingRecipe.UserID != userId {
		return nil, errs.NewForbidden("You are not allowed to update this recipe")
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
		return nil, errUpdate
	}

	return &dto.UpdateResponse{
		ID:             id,
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
	response := s.repoRecipe.Delete(id, status)
	if response != nil {
		return errs.NewInternalServerError(response.Error())
	}

	return nil
}
