package auth

import (
	"log"
	"net"

	pb "github.com/Nazerkh09/jaz/dev_microservice1/api/auth"
	"google.golang.org/grpc"
)

func RunGRPCServer() error {
	grpcServer := grpc.NewServer()

	authService := &AuthService{}

	pb.RegisterAuthServiceServer(grpcServer, authService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	log.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}
