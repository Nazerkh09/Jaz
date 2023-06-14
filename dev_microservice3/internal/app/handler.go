package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Notification represents a notification entity.
type Notification struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Message   string    `bson:"message" json:"message"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
}

// saveNotification saves the notification to the MongoDB database.
func (a *Application) saveNotification(notification Notification) error {
	collection := a.db.Database("notifications").Collection("notifications")
	_, err := collection.InsertOne(context.Background(), notification)
	if err != nil {
		return fmt.Errorf("failed to save notification: %v", err)
	}
	return nil
}

// getNotifications retrieves all notifications from the MongoDB database.
func (a *Application) getNotifications() ([]Notification, error) {
	var notifications []Notification

	collection := a.db.Database("notifications").Collection("notifications")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get notifications: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var notification Notification
		if err := cursor.Decode(&notification); err != nil {
			log.Printf("Failed to decode notification: %v", err)
			continue
		}
		notifications = append(notifications, notification)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over notifications: %v", err)
	}

	return notifications, nil
}

// sendNotificationHandler handles the request to send a notification.
func (a *Application) sendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var notification Notification
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	notification.Timestamp = time.Now()

	if err := a.saveNotification(notification); err != nil {
		http.Error(w, "Failed to save notification", http.StatusInternalServerError)
		return
	}

	if err := sendNotification(notification.Message); err != nil {
		log.Printf("Failed to send notification: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
}

// getNotificationsHandler handles the request to retrieve all notifications.
func (a *Application) getNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	notifications, err := a.getNotifications()
	if err != nil {
		http.Error(w, "Failed to get notifications", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(notifications)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// registerRoutes registers application routes and handlers.
func (a *Application) registerRoutes() {
	a.router.HandleFunc("/send-notification", a.sendNotificationHandler).Methods("POST")
	a.router.HandleFunc("/notifications", a.getNotificationsHandler).Methods("GET")
}
