package recipehandler

import (
	"strconv"

	"github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/dto"
	recipeservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/service/recipe_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/helper"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RecipeHandler struct {
	svc recipeservice.RecipeService
}

func NewRecipeHandler(r *gin.RouterGroup, svc recipeservice.RecipeService) {
	h := RecipeHandler{svc: svc}
	r.POST("/recipe/create", h.Create)
	r.GET("/recipes", h.Get)
	r.GET("/recipe/:id", h.GetById)
	r.PUT("/recipe/update/:id", h.Update)
	r.DELETE("/recipe/delete/:id", h.Delete)
}

func (h *RecipeHandler) Create(ctx *gin.Context) {
	var payload dto.CreateRequest
	idParse, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		errorFormat := errs.NewNotFound("Id not found!")
		ctx.JSON(errorFormat.StatusCode(), errorFormat)
		return
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errFormat := errs.NewUnprocessableEntity("Invalid request payload")
		ctx.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	response, errResponse := h.svc.Create(&payload, idParse, )
	if errResponse != nil {
		ctx.JSON(errResponse.StatusCode(), errResponse)
		return
	}

	helper.CreateResponse(ctx, "Successfully created recipe", response)
}

func (h *RecipeHandler) Get(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	limitStr := ctx.Query("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	recipes, err := h.svc.GetByPagination(page, limit)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	helper.OKResponse(ctx, "Successfully retrieved the data", recipes)
}

func (h *RecipeHandler) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, parseErr := uuid.Parse(idStr)
	if parseErr != nil {
		errFormat := errs.NewBadRequest("Invalid recipe ID format")
		ctx.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	recipe, err := h.svc.GetById(id)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	helper.OKResponse(ctx, "Successfully retrieved recipe detail", recipe)
}

func (h *RecipeHandler) Update(ctx *gin.Context) {
	var payload dto.UpdateRequest

	idStr := ctx.Param("id")
	id, parseErr := uuid.Parse(idStr)
	if parseErr != nil {
		errFormat := errs.NewBadRequest("Invalid recipe ID format")
		ctx.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	userID, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		errorFormat := errs.NewNotFound("User ID not found!")
		ctx.JSON(errorFormat.StatusCode(), errorFormat)
		return
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errFormat := errs.NewUnprocessableEntity("Invalid request payload")
		ctx.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	response, errResponse := h.svc.Update(id, &payload, userID)
	if errResponse != nil {
		ctx.JSON(errResponse.StatusCode(), errResponse)
		return
	}

	helper.OKResponse(ctx, "Successfully updated recipe", response)
}

func (h *RecipeHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	recipeID, err := uuid.Parse(id)
	if err != nil {
		errorFormat := errs.NewBadRequest("Invalid recipe ID")
		ctx.JSON(errorFormat.StatusCode(), errorFormat)
		return
	}

	statusStr := ctx.Query("status")
	if statusStr == "" {
		statusStr = "false"
	}

	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		errorFormat := errs.NewBadRequest("Invalid status value, must be true or false")
		ctx.JSON(errorFormat.StatusCode(), errorFormat)
		return
	}

	errDelete := h.svc.Delete(recipeID, status)
	if errDelete != nil {
		ctx.JSON(errDelete.StatusCode(), errDelete)
		return
	}

	message := "Recipe successfully deactivated"
	if status {
		message = "Recipe successfully reactivated"
	}

	helper.OKResponse(ctx, message, nil)
}
