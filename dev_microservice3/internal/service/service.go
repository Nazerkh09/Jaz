package service

import (
	"log"

	"github.com/Nazerkh09/internal/notification" // Generated Protobuf code
)

// NotificationService represents the service layer for the notification microservice.
type NotificationService struct {
	repo *repository.NotificationRepository
}

// NewNotificationService creates a new instance of NotificationService.
func NewNotificationService(repo *repository.NotificationRepository) *NotificationService {
	return &NotificationService{
		repo: repo,
	}
}

// SendNotification sends a notification and saves the notification time.
func (s *NotificationService) SendNotification(message, recipient string, timestamp int64) error {
	// Perform any necessary validation or business logic before sending the notification.

	// Create a notification request.
	req := &notification.NotificationRequest{
		Message:   message,
		Recipient: recipient,
		Timestamp: timestamp,
	}

	// Send the notification request to the gRPC server.
	// TODO: Implement gRPC client to send the notification request.

	// Save the notification time to MongoDB.
	err := s.repo.SaveNotificationTime(recipient, timestamp)
	if err != nil {
		log.Printf("Failed to save notification time: %v", err)
		return err
	}

	return nil
}
