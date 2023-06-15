package main

import (
	"encoding/json"
	"log"
	"net/http"

	"context"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/Nazerkh09/jaz/dev_microservice2/api/document"
	"github.com/Nazerkh09/jaz/dev_microservice2/internal/document"
)

func connectToMongoDBCompass() (*mongo.Client, error) {

	connectionString := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func main() {

	go document.RunGRPCServer()

	router := http.NewServeMux()

	router.HandleFunc("/api/document/create", handleCreateDocument)
	router.HandleFunc("/api/document/get", handleGetDocument)
	router.HandleFunc("/api/document/update", handleUpdateDocument)
	router.HandleFunc("/api/document/delete", handleDeleteDocument)

	log.Println("Starting HTTP/REST server on port 8081...")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func handleCreateDocument(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var createDocumentRequest pb.CreateDocumentRequest
	err := json.NewDecoder(r.Body).Decode(&createDocumentRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	documentService := document.NewDocumentService()

	// Call gRPC CreateDocument method
	createDocumentResponse, err := documentService.CreateDocument(r.Context(), &createDocumentRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(createDocumentResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleGetDocument(w http.ResponseWriter, r *http.Request) {
	// Extract document ID from URL parameter
	documentID := r.URL.Query().Get("id")

	// Create GetDocumentRequest
	getDocumentRequest := &pb.GetDocumentRequest{
		DocumentId: documentID,
	}

	documentService := document.NewDocumentService()

	// Call gRPC GetDocument method
	getDocumentResponse, err := documentService.GetDocument(r.Context(), getDocumentRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(getDocumentResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleUpdateDocument(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var updateDocumentRequest pb.UpdateDocumentRequest
	err := json.NewDecoder(r.Body).Decode(&updateDocumentRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	documentService := document.NewDocumentService()

	// Call gRPC UpdateDocument method
	updateDocumentResponse, err := documentService.UpdateDocument(r.Context(), &updateDocumentRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(updateDocumentResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleDeleteDocument(w http.ResponseWriter, r *http.Request) {
	// Create DeleteDocumentRequest
	deleteDocumentRequest := &pb.DeleteDocumentRequest{
		DocumentId: r.URL.Query().Get("id"),
	}

	documentService := document.NewDocumentService()

	// Call gRPC DeleteDocument method
	deleteDocumentResponse, err := documentService.DeleteDocument(r.Context(), deleteDocumentRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(deleteDocumentResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
