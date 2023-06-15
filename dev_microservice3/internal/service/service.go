package service

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Notification represents a notification entity.
type Notification struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Message string             `bson:"message"`
}

// Service provides methods for managing notifications.
type Service struct {
	collection *mongo.Collection
}

// NewService creates a new instance of the Service.
func NewService(collection *mongo.Collection) *Service {
	return &Service{
		collection: collection,
	}
}

// SendNotification sends a notification with the given message.
func (s *Service) SendNotification(ctx context.Context, message string) error {
	notification := Notification{
		Message: message,
	}

	_, err := s.collection.InsertOne(ctx, notification)
	if err != nil {
		return fmt.Errorf("failed to send notification: %v", err)
	}

	return nil
}

// GetNotifications retrieves all notifications.
func (s *Service) GetNotifications(ctx context.Context) ([]Notification, error) {
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get notifications: %v", err)
	}
	defer cursor.Close(ctx)

	var notifications []Notification
	for cursor.Next(ctx) {
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
