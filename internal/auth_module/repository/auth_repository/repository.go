package authrepository

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/entity"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
)

type AuthRepostiory interface {
	CreateUser(user *entity.User) (*entity.User, errs.ErrMessage)
	GetByEmail(email string) (*entity.User, errs.ErrMessage)
}
