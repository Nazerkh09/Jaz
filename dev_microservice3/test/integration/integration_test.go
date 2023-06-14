package integration

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/stretchr/testify/assert"

	"github.com/Nazerkh09/Jaz/dev_microservice3/internal/app"
	"github.com/Nazerkh09/dev_microservice3/internal/repository"
	"github.com/Nazerkh09/dev_microservice3/internal/service"
)

func TestIntegration(t *testing.T) {
	// Connect to the MongoDB test server.
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// Create a new test database and collection.
	testDB := client.Database("test_db")
	testCollection := testDB.Collection("notifications")

	// Create a new instance of the application with the test collection.
	application := app.NewApplication()
	application.SetRepository(repository.NewMongoRepository(testCollection))
	application.SetService(service.NewService(testCollection))

	// Run the application in a separate Goroutine.
	go func() {
		if err := application.Run(); err != nil {
			log.Fatalf("Application error: %v", err)
		}
	}()

	// Wait for the application to start.
	time.Sleep(1 * time.Second)

	// Run integration tests.
	t.Run("SendNotification", testSendNotification)
	t.Run("GetNotifications", testGetNotifications)
}

func testSendNotification(t *testing.T) {
	// Create a new HTTP client.
	client := &http.Client{}

	// Prepare the request data.
	data := []byte(`{"message": "Test notification"}`)

	// Send a POST request to the sendNotificationHandler endpoint.
	resp, err := client.Post("http://localhost:8000/notifications", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code.
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Check the response body.
	assert.Equal(t, "Notification sent successfully", string(body))
}

func testGetNotifications(t *testing.T) {
	// Send a GET request to the getNotificationsHandler endpoint.
	resp, err := http.Get("http://localhost:8000/notifications")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code.
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Check the response body.
	expectedResponse := `[{"message":"Test notification"}]`
	assert.Equal(t, expectedResponse, string(body))
}
