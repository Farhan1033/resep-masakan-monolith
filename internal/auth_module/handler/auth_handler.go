package authhandler

import (
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
	}

	helper.CreateResponse(ctx, "Account created successfully", response)
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

	helper.OKResponse(ctx, "Login successfully", data)
}
