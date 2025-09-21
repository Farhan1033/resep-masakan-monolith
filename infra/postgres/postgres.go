package postgressql

import (
	"fmt"
	"log"

	"github.com/Farhan1033/resep-masakan-monolith.git/infra/config"
	"github.com/Farhan1033/resep-masakan-monolith.git/internal/auth_module/entity"
	categoryentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/category_module/entity"
	ingrediententity "github.com/Farhan1033/resep-masakan-monolith.git/internal/ingredient_module/entity"
	recipeentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitPostgres() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.GetKey("DB_HOST"),
		config.GetKey("DB_USER"),
		config.GetKey("DB_PASS"),
		config.GetKey("DB_NAME"),
		config.GetKey("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	errMigrate := DB.AutoMigrate(
		entity.User{},
		categoryentity.Category{},
		ingrediententity.Ingredient{},
		recipeentity.Recipe{},
	)
	if errMigrate != nil {
		log.Fatal("Failed to migrate database")
	}

	fmt.Println("✅ Database connected successfully")
}
