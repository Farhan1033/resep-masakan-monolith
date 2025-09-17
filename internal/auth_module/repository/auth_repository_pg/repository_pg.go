package authrepositorypg

import (
	"fmt"

	"github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/entity"
	authrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/repository/auth_repository"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
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
		return nil, errs.NewInternalServerError(err.Error())
	}

	return user, nil
}

func (r *AuthRepo) GetByEmail(email string) (*entity.User, errs.ErrMessage) {
	var user entity.User

	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with this %s email not found", email))
	}

	return &user, nil
}
