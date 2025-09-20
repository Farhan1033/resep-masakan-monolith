package main

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/infra/config"
	postgressql "github.com/Farhan1033/resep-masakan-monolith.git/infra/postgres"
	authhandler "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/handler"
	authrepositorypg "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/repository/auth_repository_pg"
	authserviceimpl "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/service/auth_service_impl"
	categoryhandler "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/handler"
	categoryrepositorypg "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/repository/category_repository_pg"
	categoryserviceimp "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/service/category_service_imp"
	ingredienthandler "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/handler"
	ingredientrepositorypg "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/repository/ingredient_repository_pg"
	ingredientserviceimp "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/service/ingredient_service_imp"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	postgressql.InitPostgres()

	r := gin.Default()

	// Init Repository
	authRepo := authrepositorypg.NewAuthRepository(postgressql.DB)
	categoryRepo := categoryrepositorypg.NewCategoryRepository(postgressql.DB)
	ingredientRepo := ingredientrepositorypg.NewIngredientRepository(postgressql.DB)

	// Init Service
	authService := authserviceimpl.NewAuthService(authRepo)
	categoryService := categoryserviceimp.NewCategoryService(categoryRepo)
	ingredientService := ingredientserviceimp.NewIngredientService(ingredientRepo)

	// Setup Router
	publicGroup := r.Group("/api/v1")
	privateGroup := r.Group("/api/v1")
	privateGroup.Use(middleware.Authentication())

	// Init Handler
	authhandler.NewAuthHandler(publicGroup, authService)
	categoryhandler.NewCategoryHandler(privateGroup, categoryService)
	ingredienthandler.NewIngredientHandler(privateGroup, ingredientService)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Status": "Berhasil"})
	})

	r.Run(":8080")
}
