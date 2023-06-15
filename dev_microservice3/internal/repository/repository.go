package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Notification represents a notification entity.
type Notification struct {
	ID        string    `bson:"_id,omitempty"`
	Message   string    `bson:"message"`
	Timestamp time.Time `bson:"timestamp"`
}

// Repository provides an interface to interact with the notifications collection in MongoDB.
type Repository struct {
	collection *mongo.Collection
}

// NewRepository creates a new instance of the Repository.
func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{
		collection: collection,
	}
}

// SaveNotification saves the notification to the MongoDB collection.
func (r *Repository) SaveNotification(notification Notification) error {
	_, err := r.collection.InsertOne(context.Background(), notification)
	if err != nil {
		return fmt.Errorf("failed to save notification: %v", err)
	}
	return nil
}

// GetNotifications retrieves all notifications from the MongoDB collection.
func (r *Repository) GetNotifications() ([]Notification, error) {
	var notifications []Notification

	cursor, err := r.collection.Find(context.Background(), bson.M{})
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
