package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/your-username/your-app/internal/service"
)

func TestUnit(t *testing.T) {
	// Create a new instance of the service.
	svc := service.NewService()

	// Run unit tests.
	t.Run("SaveNotification", testSaveNotification(svc))
	t.Run("GetNotifications", testGetNotifications(svc))
}

func testSaveNotification(svc *service.Service) func(t *testing.T) {
	return func(t *testing.T) {
		// Create a new notification.
		notification := &service.Notification{
			Message: "Test notification",
		}

		// Save the notification.
		err := svc.SaveNotification(notification)
		assert.NoError(t, err)
		assert.NotEmpty(t, notification.ID)

		// Verify that the notification was saved correctly.
		savedNotification, err := svc.GetNotification(notification.ID)
		assert.NoError(t, err)
		assert.Equal(t, notification.Message, savedNotification.Message)
	}
}

func testGetNotifications(svc *service.Service) func(t *testing.T) {
	return func(t *testing.T) {
		// Get the list of notifications.
		notifications, err := svc.GetNotifications()
		assert.NoError(t, err)
		assert.Empty(t, notifications)
	}
}
