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
	var output *os.File
	var err error

	if err = os.MkdirAll("/tmp/logs", 0755); err != nil {
		fmt.Println("Cannot create /tmp/logs, fallback to stdout")
		output = os.Stdout
	} else {
		output, err = os.OpenFile("/tmp/logs/event.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Cannot open /tmp/logs/event.log, fallback to stdout")
			output = os.Stdout
		}
	}

	eventLogger = log.New(output, "[EVENT] ", log.Ldate|log.Ltime)
	log.SetOutput(output)
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
