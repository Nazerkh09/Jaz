package document

import (
	"context"
	"testing"

	pb "github.com/Nazerkh09/jaz/dev_microservice2/api/document"
	"github.com/stretchr/testify/assert"
)

func TestDocumentService_CreateDocument(t *testing.T) {
	service := &DocumentService{}

	req := &pb.CreateDocumentRequest{
		Title:   "Sample Document",
		Content: "Sample content of the document",
	}

	res, err := service.CreateDocument(context.Background(), req)

	assert.Nil(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "Document created successfully", res.Message)
}

// Implement tests for other methods
