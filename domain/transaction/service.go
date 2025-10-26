package transaction

import (
	"context"
	"errors"

	"sandbox/application/dto"
)

// Service provides domain business logic for transactions
type Service struct {
	extractor ExtractorRepository
}

// NewService creates a new transaction service
func NewService(extractor ExtractorRepository) *Service {
	return &Service{
		extractor: extractor,
	}
}

// ExtractTransactions extracts transactions from multiple documents
func (s *Service) ExtractTransactions(ctx context.Context, documents []Document) (*dto.RecapReportDTO, error) {
	if len(documents) == 0 {
		return nil, errors.New("no documents provided")
	}

	recapReport, err := s.extractor.ExtractFromDocuments(ctx, documents)
	if err != nil {
		return nil, err
	}

	// Additional business logic can be added here
	// For example: validation, deduplication, etc.

	return recapReport, nil
}
