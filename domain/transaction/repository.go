package transaction

import (
	"context"

	"sandbox/application/dto"
)

// ExtractorRepository defines the contract for extracting transactions from documents
type ExtractorRepository interface {
	ExtractFromDocuments(ctx context.Context, documents []Document) (*dto.RecapReportDTO, error)
}

// Document represents a document to be processed
type Document struct {
	Content  []byte
	MimeType string
	Filename string
}
