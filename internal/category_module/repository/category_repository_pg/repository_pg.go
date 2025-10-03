package categoryrepositorypg

import (
	"fmt"

	categoryentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/entity"
	categoryrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/repository/category_repository"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) categoryrepository.CategoryRepository {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) Create(category *categoryentity.Category) (*categoryentity.Category, errs.ErrMessage) {
	if err := r.db.Create(category).Error; err != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to create category: %v", err))
	}
	return category, nil
}

func (r *CategoryRepo) Get() ([]categoryentity.Category, errs.ErrMessage) {
	var categories []categoryentity.Category
	if err := r.db.Where("is_active = true").Order("name ASC").Find(&categories).Error; err != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to get categories: %v", err))
	}
	return categories, nil
}

func (r *CategoryRepo) GetByName(name string) (*categoryentity.Category, errs.ErrMessage) {
	var category categoryentity.Category
	if err := r.db.Where("name = ?", name).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFound(fmt.Sprintf("category with name '%s' not found", name))
		}
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to get category by name: %v", err))
	}
	return &category, nil
}

func (r *CategoryRepo) GetById(id uuid.UUID) (*categoryentity.Category, errs.ErrMessage) {
	var category categoryentity.Category
	if err := r.db.Where("id = ? AND is_active = true", id).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFound(fmt.Sprintf("category with id '%s' not found or inactive", id))
		}
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to get category by id: %v", err))
	}
	return &category, nil
}

func (r *CategoryRepo) Delete(id uuid.UUID, status bool) errs.ErrMessage {
	result := r.db.Model(&categoryentity.Category{}).Where("id = ?", id).Update("is_active", status)
	if result.Error != nil {
		return errs.NewInternalServerError(fmt.Sprintf("failed to update category status: %v", result.Error))
	}
	if result.RowsAffected == 0 {
		return errs.NewNotFound(fmt.Sprintf("category with id '%s' not found", id))
	}
	return nil
}
