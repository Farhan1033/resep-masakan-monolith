package detailrecipehandler

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/dto"
	detailrecipeservice "github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/service/detail_recipe_service"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/helper"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DetailRecipeHandler struct {
	svc detailrecipeservice.DetailRecipeService
}

func NewDetailRecipeHandler(r *gin.RouterGroup, svc detailrecipeservice.DetailRecipeService) {
	h := DetailRecipeHandler{svc: svc}
	r.POST("/recipe-ingredient/create", h.CreateDetailRecipe)
	r.GET("/recipe-ingredients", h.GetAllDetailRecipes)
	r.GET("/recipe-ingredient/:id", h.GetDetailRecipeById)
	r.GET("/ingredient/recipe/:recipe_id", h.GetDetailRecipeByRecipeId)
	r.PUT("/recipe-ingredient/update/:id", h.UpdateDetailRecipe)
	r.DELETE("/recipe-ingredient/delete/:id", h.DeleteDetailRecipe)
}

func (h *DetailRecipeHandler) CreateDetailRecipe(c *gin.Context) {
	var req dto.CreateDetailRecipeRequest
	idParse, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		errorFormat := errs.NewNotFound("Unauthorized user!")
		c.JSON(errorFormat.StatusCode(), errorFormat)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errFormat := errs.NewUnprocessableEntity("Invalid request payload")
		c.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	result, errMsg := h.svc.CreateDetailRecipe(idParse, &req)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode(), errMsg)
		return
	}

	helper.CreateResponse(c, "Detail recipe created successfully", result)
}

func (h *DetailRecipeHandler) GetAllDetailRecipes(c *gin.Context) {
	result, errMsg := h.svc.GetAllDetailRecipes()
	if errMsg != nil {
		c.JSON(errMsg.StatusCode(), errMsg)
		return
	}

	helper.OKResponse(c, "Detail recipes retrieved successfully", result)
}

func (h *DetailRecipeHandler) GetDetailRecipeById(c *gin.Context) {
	idStr := c.Param("id")

	id, parseErr := uuid.Parse(idStr)
	if parseErr != nil {
		errFormat := errs.NewBadRequest("Invalid recipe ID format")
		c.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	result, errMsg := h.svc.GetDetailRecipeById(id)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode(), errMsg)
		return
	}

	helper.OKResponse(c, "Detail recipes retrieved successfully", result)
}

func (h *DetailRecipeHandler) GetDetailRecipeByRecipeId(c *gin.Context) {
	recipeId := c.Param("recipe_id")

	idRecipe, parseErr := uuid.Parse(recipeId)
	if parseErr != nil {
		errFormat := errs.NewBadRequest("Invalid recipe ID format")
		c.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	result, errMsg := h.svc.GetDetailRecipeByRecipeId(idRecipe)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode(), errMsg)
		return
	}

	helper.OKResponse(c, "Detail recipes retrieved successfully", result)
}

func (h *DetailRecipeHandler) UpdateDetailRecipe(c *gin.Context) {
	idStr := c.Param("id")

	id, parseErr := uuid.Parse(idStr)
	if parseErr != nil {
		errFormat := errs.NewBadRequest("Invalid recipe ID format")
		c.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	var req dto.UpdateDetailRecipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errFormat := errs.NewUnprocessableEntity("Invalid request payload")
		c.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	errMsg := h.svc.UpdateDetailRecipe(id, &req)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode(), errMsg)
		return
	}

	helper.OKResponse(c, "Detail recipe updated successfully", nil)
}

func (h *DetailRecipeHandler) DeleteDetailRecipe(c *gin.Context) {
	idStr := c.Param("id")

	id, parseErr := uuid.Parse(idStr)
	if parseErr != nil {
		errFormat := errs.NewBadRequest("Invalid recipe ID format")
		c.JSON(errFormat.StatusCode(), errFormat)
		return
	}

	errMsg := h.svc.DeleteDetailRecipe(id)
	if errMsg != nil {
		c.JSON(errMsg.StatusCode(), errMsg)
		return
	}

	helper.OKResponse(c, "Detail recipe deleted successfully", nil)
}
