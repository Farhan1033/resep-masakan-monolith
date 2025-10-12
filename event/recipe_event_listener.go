package event

import (
	"fmt"
	"log"
	"os"
	"time"

	recipeentity "github.com/Farhan1033/resep-masakan-monolith.git/internal/recipe_module/entity"
)

var eventLogger *log.Logger

func InitEventLogger() {
	file, err := os.OpenFile("logs/event.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open event log file: %v", err)
	}

	eventLogger = log.New(file, "[EVENT] ", log.Ldate|log.Ltime)
}

func RegisterRecipeEvents(bus *EventBus) {
	bus.Subscribe("recipeCreated", func(data interface{}) {
		if recipe, ok := data.(*recipeentity.Recipe); ok {
			message := fmt.Sprintf("Recipe created: %s | created at: %s",
				recipe.Title, time.Now().Format(time.RFC3339))
			fmt.Println(message)
			eventLogger.Println(message)
		}
	})

	bus.Subscribe("recipeUpdated", func(data interface{}) {
		if recipe, ok := data.(*recipeentity.Recipe); ok {
			message := fmt.Sprintf("Recipe updated: %s | updated at: %s",
				recipe.Title, time.Now().Format(time.RFC3339))
			fmt.Println(message)
			eventLogger.Println(message)
		}
	})

	bus.Subscribe("recipeDeleted", func(data interface{}) {
		if recipe, ok := data.(*recipeentity.Recipe); ok {
			message := fmt.Sprintf("Recipe deleted: %s | deleted at: %s",
				recipe.Title, time.Now().Format(time.RFC3339))
			fmt.Println(message)
			eventLogger.Println(message)
		}
	})
}
