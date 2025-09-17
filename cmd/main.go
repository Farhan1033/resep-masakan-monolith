package main

import (
	"github.com/Farhan1033/resep-masakan-monolith.git/infra/config"
	postgressql "github.com/Farhan1033/resep-masakan-monolith.git/infra/postgres"
	authhandler "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/handler"
	authrepositorypg "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/repository/auth_repository_pg"
	authserviceimpl "github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/service/auth_service_impl"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	postgressql.InitPostgres()

	r := gin.Default()

	// Init Repository
	authRepo := authrepositorypg.NewAuthRepository(postgressql.DB)

	// Init Service
	authService := authserviceimpl.NewAuthService(authRepo)

	// Setup Router
	publicGroup := r.Group("/api/v1")

	// Init Handler
	authhandler.NewAuthHandler(publicGroup, authService)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Status": "Berhasil"})
	})

	r.Run(":8080")
}
