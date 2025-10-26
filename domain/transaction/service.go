package transaction

import (
	"context"
	"errors"
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
func (s *Service) ExtractTransactions(ctx context.Context, documents []Document) ([]*Transaction, error) {
	if len(documents) == 0 {
		return nil, errors.New("no documents provided")
	}

	transactions, err := s.extractor.ExtractFromDocuments(ctx, documents)
	if err != nil {
		return nil, err
	}

	// Additional business logic can be added here
	// For example: validation, deduplication, etc.

	return transactions, nil
}
