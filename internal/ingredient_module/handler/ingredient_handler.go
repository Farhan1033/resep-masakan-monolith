package ingredienthandler

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/dto"
	ingredientservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/service/ingredient_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IngredientHandler struct {
	svc ingredientservice.IngredientService
}

func NewIngredientHandler(r *gin.RouterGroup, svc ingredientservice.IngredientService) {
	h := IngredientHandler{svc: svc}
	r.POST("/ingredient/create", h.Create)
	r.GET("/ingredients", h.Get)
	r.PATCH("/ingredient/delete", h.Delete)
}

func (h *IngredientHandler) Create(ctx *gin.Context) {
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

	helper.CreateResponse(ctx, "Successfully created category", response)
}

func (h *IngredientHandler) Get(ctx *gin.Context) {
	response, err := h.svc.Get()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	helper.OKResponse(ctx, "Success", response)
}

func (h *IngredientHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Query("id")
	if idParam == "" {
		errFormat := errs.NewBadRequest("Invalid id category")
		ctx.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	idString, _ := uuid.Parse(idParam)
	resonse := h.svc.Delete(idString)
	if resonse != nil {
		ctx.JSON(resonse.StatusCode(), resonse)
		return
	}

	helper.OKResponse(ctx, "Successfully deleted data", nil)
}
