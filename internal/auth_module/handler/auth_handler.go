package authhandler

import (
	"strings"

	"github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/dto"
	authservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/service/auth_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/helper"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc authservice.AuthService
}

func NewAuthHandler(r *gin.RouterGroup, svc authservice.AuthService) {
	h := AuthHandler{svc: svc}
	r.POST("/auth/register", h.Register)
	r.POST("/auth/login", h.Login)
	r.POST("/auth/logout", h.Logout)
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var payload dto.CreateRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errFormat := errs.NewUnprocessableEntity("Invalid request payload")
		ctx.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	response, err := h.svc.Create(&payload)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	helper.CreateResponse(ctx, "User registered successfully", response)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var payload dto.LoginRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		newError := errs.NewUnprocessableEntity("Invalid request payload")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	data, err := h.svc.Login(&payload)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	helper.OKResponse(ctx, "Login successful", data)
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		errsFormat := errs.NewUnauthorized("Authorization header is required")
		ctx.JSON(errsFormat.StatusCode(), errsFormat)
		ctx.Abort()
		return
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		errsFormat := errs.NewUnauthorized("Invalid authorization header format")
		ctx.JSON(errsFormat.StatusCode(), errsFormat)
		ctx.Abort()
		return
	}

	helper.OKResponse(ctx, "Logged out successfully", nil)
}
