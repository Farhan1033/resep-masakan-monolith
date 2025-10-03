package main

import (
	"log"
	"os"
	"time"

	"github.com/Farhan1033/resep-masakan-monolith.git/infra/config"
	postgressql "github.com/Farhan1033/resep-masakan-monolith.git/infra/postgres"
	authhandler "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/handler"
	authrepositorypg "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/repository/auth_repository_pg"
	authserviceimpl "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/service/auth_service_impl"
	categoryhandler "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/handler"
	categoryrepositorypg "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/repository/category_repository_pg"
	categoryserviceimp "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/service/category_service_imp"
	detailrecipehandler "github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/handler"
	detailreciperepositorypg "github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/repository/detail_recipe_repository_pg"
	detailrecipeserviceimpl "github.com/Farhan1033/resep-masakan-monolith.git/internal/detail_recipe_module/service/detail_recipe_service_impl"
	ingredienthandler "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/handler"
	ingredientrepositorypg "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/repository/ingredient_repository_pg"
	ingredientserviceimp "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/service/ingredient_service_imp"
	recipehandler "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/handler"
	reciperepositorypg "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/repository/recipe_repository_pg"
	recipeserviceimp "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/service/recipe_service_imp"
	stephandler "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/handler"
	steprepositorypg "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/repository/step_repository_pg"
	stepserviceimpl "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_steps_module/service/step_service_impl"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// Load ENV
	config.LoadEnv()
	infoLogger.Println("Environment variables loaded")

	// Init DB
	postgressql.InitPostgres()
	infoLogger.Println("PostgreSQL connected successfully")

	// Setup Gin
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	log.Println("Gin mode: RELEASE")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Init Repository
	authRepo := authrepositorypg.NewAuthRepository(postgressql.DB)
	categoryRepo := categoryrepositorypg.NewCategoryRepository(postgressql.DB)
	ingredientRepo := ingredientrepositorypg.NewIngredientRepository(postgressql.DB)
	detailRepo := detailreciperepositorypg.NewDetailRecipeRepository(postgressql.DB)
	stepRepo := steprepositorypg.NewRecipeStepRepository(postgressql.DB)
	recipeRepo := reciperepositorypg.NewRecipeRepository(postgressql.DB)

	// Init Service
	authService := authserviceimpl.NewAuthService(authRepo)
	categoryService := categoryserviceimp.NewCategoryService(categoryRepo)
	ingredientService := ingredientserviceimp.NewIngredientService(ingredientRepo)
	detailService := detailrecipeserviceimpl.NewDetailRecipeService(detailRepo, authRepo)
	stepService := stepserviceimpl.NewRecipeStepService(stepRepo)
	recipeService := recipeserviceimp.NewRecipeService(
		recipeRepo,
		authRepo,
		categoryRepo,
		stepService,
		detailService,
	)

	// Setup Router
	publicGroup := r.Group("/api/v1")
	privateGroup := r.Group("/api/v1")
	privateGroup.Use(middleware.Authentication())

	// Init Handler
	authhandler.NewAuthHandler(publicGroup, authService)
	categoryhandler.NewCategoryHandler(privateGroup, categoryService)
	ingredienthandler.NewIngredientHandler(privateGroup, ingredientService)
	detailrecipehandler.NewDetailRecipeHandler(privateGroup, detailService)
	stephandler.NewRecipeStepHandler(privateGroup, stepService)
	recipehandler.NewRecipeHandler(privateGroup, recipeService)

	// Health Check
	r.GET("/kaithheathcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "Recipe App is running smoothly 🚀",
		})
	})

	infoLogger.Println("Server running")
	if err := r.Run("0.0.0.0:" + config.GetKey("APP_PORT")); err != nil {
		errorLogger.Fatalf("Failed to start server: %v", err)
	}
}
