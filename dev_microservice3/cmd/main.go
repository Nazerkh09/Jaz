package main

import (
	"log"

	"your-app-name/internal/app"
)

func main() {
	// Initialize the application.
	application := app.NewApplication()

	// Start the application.
	if err := application.Run(); err != nil {
		log.Fatalf("Failed to start the application: %v", err)
	}
}
