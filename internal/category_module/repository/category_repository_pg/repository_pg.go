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

func (r *CategoryRepo) Create(catgeory *categoryentity.Category) (*categoryentity.Category, errs.ErrMessage) {
	if err := r.db.Create(catgeory).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return catgeory, nil
}

func (r *CategoryRepo) Get() ([]categoryentity.Category, errs.ErrMessage) {
	var categories []categoryentity.Category

	if err := r.db.Where("is_active = true").Order("name ASC").Find(&categories).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return categories, nil
}

func (r *CategoryRepo) GetByName(name string) (*categoryentity.Category, errs.ErrMessage) {
	var category categoryentity.Category

	if err := r.db.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, errs.NewFound(fmt.Sprintf("This %s category already exists", name))
	}

	return &category, nil
}

func (r *CategoryRepo) GetById(id uuid.UUID) (*categoryentity.Category, errs.ErrMessage) {
	var category categoryentity.Category

	if err := r.db.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, errs.NewNotFound("This category not found or inactive")
	}
	return &category, nil
}

func (r *CategoryRepo) Delete(id uuid.UUID, status bool) errs.ErrMessage {
	if err := r.db.Model(&categoryentity.Category{}).Where("id = ?", id).Update("is_active", status).Error; err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}
