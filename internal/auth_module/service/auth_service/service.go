package authservice

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/dto"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
)

type AuthService interface {
	Create(payload *dto.CreateRequest) (*dto.CreateAuthResponse, errs.ErrMessage)
	Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.ErrMessage)
}
