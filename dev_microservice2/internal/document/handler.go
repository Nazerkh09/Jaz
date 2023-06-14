package document

import (
	"log"
	"net"

	pb "github.com/Nazerkh09/jaz/dev_microservice2/api/document"
	"google.golang.org/grpc"
)

func RunGRPCServer() error {
	grpcServer := grpc.NewServer()

	documentService := &DocumentService{}

	pb.RegisterDocumentServiceServer(grpcServer, documentService)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	log.Println("Starting gRPC server on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}
