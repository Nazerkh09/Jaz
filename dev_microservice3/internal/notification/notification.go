package notification

import (
	"context"
	"log"
	"net"

	pb "your-app-name/internal/notification" // Generated Protobuf code

	"google.golang.org/grpc"
)

// NotificationServer represents the gRPC server for the notification service.
type NotificationServer struct {
}

// SendNotification handles the SendNotification gRPC request.
func (s *NotificationServer) SendNotification(ctx context.Context, req *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	// Handle the notification request, save the time to MongoDB, and perform any necessary logic.
	log.Printf("Received notification request: %v", req)

	// Save the notification time to MongoDB using the repository.

	// Return the response.
	return &pb.NotificationResponse{
		Success: true,
		Message: "Notification sent successfully",
	}, nil
}

// RunServer starts the gRPC server for the notification service.
func RunServer() {
	// Create a listener on TCP port 50051.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server.
	s := grpc.NewServer()

	// Register the notification service server.
	pb.RegisterNotificationServiceServer(s, &NotificationServer{})

	// Start the gRPC server.
	log.Println("Starting gRPC server on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
