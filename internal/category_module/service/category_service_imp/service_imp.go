package categoryserviceimp

import (
	"fmt"

	"github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/dto"
	categoryentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/entity"
	categoryrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/repository/category_repository"
	categoryservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/service/category_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CategorySvc struct {
	repo     categoryrepository.CategoryRepository
	validate *validator.Validate
}

func NewCategoryService(repo categoryrepository.CategoryRepository) categoryservice.CategoryService {
	return &CategorySvc{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *CategorySvc) Create(payload *dto.CreateCategoryRequest) (*dto.CategoryResponse, errs.ErrMessage) {
	if err := s.validate.Struct(payload); err != nil {
		formatError := validation.FormatValidationError(err)
		return nil, errs.NewBadRequest(fmt.Sprintf("Required: %s", formatError))
	}

	if _, err := s.repo.GetByName(payload.Name); err == nil {
		return nil, errs.NewFound("This category already exists")
	}

	craeteNew := &categoryentity.Category{
		Name: payload.Name,
	}

	createCategory, err := s.repo.Create(craeteNew)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	response := &dto.CategoryResponse{
		ID:        createCategory.ID,
		Name:      createCategory.Name,
		IsActive:  createCategory.IsActive,
		CreatedAt: createCategory.CreatedAt,
	}

	return response, nil
}

func (s *CategorySvc) Get() ([]*dto.CategoryResponse, errs.ErrMessage) {
	categories, err := s.repo.Get()
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	if len(categories) == 0 {
		return nil, errs.NewNotFound("Data not found!")
	}

	response := make([]*dto.CategoryResponse, len(categories))
	for i, catgeory := range categories {
		response[i] = &dto.CategoryResponse{
			ID:        catgeory.ID,
			Name:      catgeory.Name,
			IsActive:  catgeory.IsActive,
			CreatedAt: catgeory.CreatedAt,
			UpdatedAt: catgeory.UpdatedAt,
		}
	}

	return response, nil
}

func (s *CategorySvc) Delete(id uuid.UUID) errs.ErrMessage {
	if _, err := s.repo.GetById(id); err != nil {
		return errs.NewFound(err.Error())
	}

	response := s.repo.Delete(id, false)
	if response != nil {
		return errs.NewInternalServerError(response.Error())
	}

	return nil
}
