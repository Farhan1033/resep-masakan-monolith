package categoryhandler

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/dto"
	categoryservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/service/category_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	svc categoryservice.CategoryService
}

func NewCategoryHandler(r *gin.RouterGroup, svc categoryservice.CategoryService) {
	h := CategoryHandler{svc: svc}
	r.POST("/category/create", h.Create)
	r.GET("/categories", h.Get)
	r.DELETE("/category/delete/:id", h.Delete)
}

func (h *CategoryHandler) Create(ctx *gin.Context) {
	var payload dto.CreateCategoryRequest

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

func (h *CategoryHandler) Get(ctx *gin.Context) {
	response, err := h.svc.Get()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	helper.OKResponse(ctx, "Success", response)
}

func (h *CategoryHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
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
