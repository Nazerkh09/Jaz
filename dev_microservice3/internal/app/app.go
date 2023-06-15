package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Application represents the main application.
type Application struct {
	router *mux.Router
	db     *mongo.Client
}

// NewApplication creates a new instance of the Application.
func NewApplication() *Application {
	// Initialize the MongoDB client.
	db, err := connectToMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create a new Gorilla mux router.
	router := mux.NewRouter()

	return &Application{
		router: router,
		db:     db,
	}
}

// Run starts the application.
func (a *Application) Run() error {
	// Register application routes and handlers.
	a.registerRoutes()

	// Start the HTTP server.
	return a.startServer()
}

// connectToMongoDB initializes the MongoDB client and returns a connection.
func connectToMongoDB() (*mongo.Client, error) {
	// Set up MongoDB connection options.
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB to check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")

	return client, nil
}

// registerRoutes registers application routes and handlers.
func (a *Application) registerRoutes() {
	a.router.HandleFunc("/send-notification", a.sendNotificationHandler).Methods("POST")
}

// sendNotificationHandler handles the request to send a notification.
func (a *Application) sendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	notification := "This is a notification message"
	err := sendNotification(notification)
	if err != nil {
		http.Error(w, "Failed to send notification", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Notification sent successfully"))
}

// sendNotification sends the notification.
func sendNotification(notification string) error {
	log.Printf("Sending notification: %s", notification)
	// TODO: Implement the code to send the notification

	return nil
}

// startServer starts the HTTP server.
func (a *Application) startServer() error {
	serverAddr := ":8080"
	log.Printf("Starting server at %s", serverAddr)
	return http.ListenAndServe(serverAddr, a.router)
}
