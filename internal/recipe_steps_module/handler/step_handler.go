package stephandler

import (
	"strconv"

	stepdto "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/dto"
	stepservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/service/step_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/helper"
	"github.com/gin-gonic/gin"
)

type RecipeStepHandler struct {
	svc stepservice.RecipeStepService
}

func NewRecipeStepHandler(r *gin.RouterGroup, svc stepservice.RecipeStepService) {
	h := RecipeStepHandler{svc: svc}
	r.POST("/step/create", h.Create)
	r.GET("/recipe/step", h.Get)
	r.PUT("/step/update/:id", h.Update)
	r.DELETE("/step/delete/:id", h.Delete)
}

func (h *RecipeStepHandler) Create(c *gin.Context) {
	var payload stepdto.CreateRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		errFormat := errs.NewUnprocessableEntity("Invalid request payload")
		c.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	response, err := h.svc.Create(&payload)
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	helper.CreateResponse(c, "Successfully created category", response)
}

func (h *RecipeStepHandler) Get(c *gin.Context) {
	response, err := h.svc.Get()
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	helper.OKResponse(c, "Success", response)
}

func (h *RecipeStepHandler) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		errFormat := errs.NewBadRequest("Invalid id category")
		ctx.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	var payload stepdto.UpdateRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errFormat := errs.NewBadRequest("Invalid request body: " + err.Error())
		ctx.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	idString, _ := strconv.Atoi(idParam)
	resonse := h.svc.Update(uint(idString), &payload)
	if resonse != nil {
		ctx.JSON(resonse.StatusCode(), resonse)
		return
	}

	helper.OKResponse(ctx, "Successfully updated data", nil)
}

func (h *RecipeStepHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		errFormat := errs.NewBadRequest("Invalid id category")
		ctx.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	idString, _ := strconv.Atoi(idParam)
	resonse := h.svc.Delete(uint(idString))
	if resonse != nil {
		ctx.JSON(resonse.StatusCode(), resonse)
		return
	}

	helper.OKResponse(ctx, "Successfully deleted data", nil)
}
