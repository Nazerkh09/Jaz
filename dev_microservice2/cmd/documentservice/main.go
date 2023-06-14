package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Nazerkh09/jaz/dev_microservice2/internal/document"
)

func main() {
	// Start the document service
	go document.RunGRPCServer()

	router := http.NewServeMux()

	router.HandleFunc("/api/document/create", handleCreateDocument)
	router.HandleFunc("/api/document/:id", handleGetDocument)
	router.HandleFunc("/api/document/:id/update", handleUpdateDocument)
	router.HandleFunc("/api/document/:id/delete", handleDeleteDocument)

	log.Println("Starting HTTP/REST server on port 8081...")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func handleCreateDocument(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var createDocumentRequest document.CreateDocumentRequest
	err := json.NewDecoder(r.Body).Decode(&createDocumentRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Perform document creation logic
	newDocument, err := document.CreateDocument(createDocumentRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return created document in the response
	response := struct {
		Document document.Document `json:"document"`
	}{
		Document: newDocument,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleGetDocument(w http.ResponseWriter, r *http.Request) {
	// Extract document ID from URL parameter
	documentID := r.URL.Query().Get("id")

	// Perform document retrieval logic
	document, err := document.GetDocument(documentID)
	if err != nil {
		// Handle error, such as document not found
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Return the document in the response
	response := struct {
		Document document.Document `json:"document"`
	}{
		Document: document,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleUpdateDocument(w http.ResponseWriter, r *http.Request) {
	// Extract document ID from URL parameter
	documentID := r.URL.Query().Get("id")

	// Parse request body
	var updateDocumentRequest document.UpdateDocumentRequest
	err := json.NewDecoder(r.Body).Decode(&updateDocumentRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Perform document update logic
	updatedDocument, err := document.UpdateDocument(documentID, updateDocumentRequest)
	if err != nil {
		// Handle error, such as document not found or validation error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the updated document in the response
	response := struct {
		Document document.Document `json:"document"`
	}{
		Document: updatedDocument,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleDeleteDocument(w http.ResponseWriter, r *http.Request) {
	// Extract document ID from URL parameter
	documentID := r.URL.Query().Get("id")

	// Perform document deletion logic
	err := document.DeleteDocument(documentID)
	if err != nil {
		// Handle error, such as document not found or deletion failed
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return success response
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Document deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
