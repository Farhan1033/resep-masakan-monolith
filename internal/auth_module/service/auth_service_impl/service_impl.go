package authserviceimpl

import (
	"fmt"

	"github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/dto"
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/entity"
	authrepository "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/repository/auth_repository"
	authservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/service/auth_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/middleware"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/validation"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AuthSvc struct {
	repo     authrepository.AuthRepostiory
	validate *validator.Validate
}

func NewAuthService(repo authrepository.AuthRepostiory) authservice.AuthService {
	return &AuthSvc{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *AuthSvc) Create(payload *dto.CreateRequest) (*dto.CreateAuthResponse, errs.ErrMessage) {
	if err := s.validate.Struct(payload); err != nil {
		formatedValidation := validation.FormatValidationError(err)
		return nil, errs.NewBadRequest(fmt.Sprintf("Validation failed: %s", formatedValidation))
	}

	_, err := s.repo.GetByEmail(payload.Email)
	if err == nil {
		return nil, errs.NewBadRequest("Email already exist!")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(hash)

	userPayload := &entity.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	registeredUser, err := s.repo.CreateUser(userPayload)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	response := &dto.CreateAuthResponse{
		ID:    registeredUser.ID,
		Name:  registeredUser.Name,
		Email: registeredUser.Email,
	}

	return response, nil
}

func (s *AuthSvc) Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.ErrMessage) {
	if err := s.validate.Struct(payload); err != nil {
		formatedValidation := validation.FormatValidationError(err)
		return nil, errs.NewBadRequest(fmt.Sprintf("Validation failed: %s", formatedValidation))
	}

	existingUser, err := s.repo.GetByEmail(payload.Email)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	if existingUser == nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with this email %s not found!", payload.Email))
	}

	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(payload.Password)) != nil {
		return nil, errs.NewBadRequest("Wrong password")
	}

	token, err := middleware.CreateToken(existingUser.ID, existingUser.Email)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	response := &dto.LoginResponse{
		Message: "Success login",
		Token:   token,
	}

	return response, nil
}
