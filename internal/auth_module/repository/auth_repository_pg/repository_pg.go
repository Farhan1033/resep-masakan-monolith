package authrepositorypg

import (
	"errors"
	"fmt"

	"github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/entity"
	authrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/repository/auth_repository"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authrepository.AuthRepostiory {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) CreateUser(user *entity.User) (*entity.User, errs.ErrMessage) {
	if err := r.db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errs.NewBadRequest("user with this email already exists")
		}
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to create user: %v", err))
	}
	return user, nil
}

func (r *AuthRepo) GetByEmail(email string) (*entity.User, errs.ErrMessage) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFound(fmt.Sprintf("user with email %s not found", email))
		}
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to get user by email: %v", err))
	}
	return &user, nil
}

func (r *AuthRepo) GetById(id uuid.UUID) (*entity.User, errs.ErrMessage) {
	var user entity.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFound(fmt.Sprintf("user with id %s not found", id))
		}
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to get user by id: %v", err))
	}
	return &user, nil
}
