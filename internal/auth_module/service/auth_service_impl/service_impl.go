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
		return nil, errs.NewBadRequest(fmt.Sprintf("validation failed: %s", formatedValidation))
	}

	existingUser, err := s.repo.GetByEmail(payload.Email)
	if err != nil {
		if err.StatusCode() != 404 {
			return nil, errs.NewInternalServerError(fmt.Sprintf("failed to check existing user: %s", err.Message()))
		}
	} else if existingUser != nil {
		return nil, errs.NewBadRequest("user with this email already exists")
	}

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to hash password: %v", hashErr))
	}

	userPayload := &entity.User{
		Email:    payload.Email,
		UserName: payload.UserName,
		FullName: payload.FullName,
		Password: string(hashedPassword),
	}

	registeredUser, err := s.repo.CreateUser(userPayload)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateAuthResponse{
		ID:        registeredUser.ID,
		Email:     registeredUser.Email,
		UserName:  registeredUser.UserName,
		FullName:  registeredUser.FullName,
		CreatedAt: registeredUser.CreatedAt.Local().String(),
	}

	return response, nil
}

func (s *AuthSvc) Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.ErrMessage) {
	if err := s.validate.Struct(payload); err != nil {
		formatedValidation := validation.FormatValidationError(err)
		return nil, errs.NewBadRequest(fmt.Sprintf("validation failed: %s", formatedValidation))
	}

	existingUser, err := s.repo.GetByEmail(payload.Email)
	if err != nil {
		if err.StatusCode() == 404 {
			return nil, errs.NewUnauthorized("invalid email or password")
		}
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to get user: %s", err.Message()))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(payload.Password)); err != nil {
		return nil, errs.NewUnauthorized("invalid email or password")
	}

	token, tokenErr := middleware.CreateToken(existingUser.ID, existingUser.Email)
	if tokenErr != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("failed to generate token: %s", tokenErr.Message()))
	}

	data := &dto.LoginDataResponse{
		ID:       existingUser.ID,
		Email:    existingUser.Email,
		UserName: existingUser.UserName,
		FullName: existingUser.FullName,
	}

	response := &dto.LoginResponse{
		Data:  data,
		Token: token,
	}

	return response, nil
}
