package main

import (
	"log"
)

func main() {
	// Initialize the application.
	application := app.NewApplication()

	// Start the application.
	if err := application.Run(); err != nil {
		log.Fatalf("Failed to start the application: %v", err)
	}
}
