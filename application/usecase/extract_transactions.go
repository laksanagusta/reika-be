package usecase

import (
	"context"

	"sandbox/application/dto"
	"sandbox/domain/transaction"
)

// ExtractTransactionsUseCase handles the business use case for extracting transactions
type ExtractTransactionsUseCase struct {
	transactionService *transaction.Service
}

// NewExtractTransactionsUseCase creates a new use case instance
func NewExtractTransactionsUseCase(transactionService *transaction.Service) *ExtractTransactionsUseCase {
	return &ExtractTransactionsUseCase{
		transactionService: transactionService,
	}
}

// Execute performs the transaction extraction use case
func (uc *ExtractTransactionsUseCase) Execute(ctx context.Context, req dto.ExtractTransactionsRequest) (*dto.ExtractTransactionsResponse, error) {
	// Convert DTOs to domain objects
	documents := make([]transaction.Document, len(req.Files))
	for i, file := range req.Files {
		documents[i] = transaction.Document{
			Content:  file.Content,
			MimeType: file.MimeType,
			Filename: file.Filename,
		}
	}

	// Execute domain logic
	recapReport, err := uc.transactionService.ExtractTransactions(ctx, documents)
	if err != nil {
		return nil, err
	}

	return &dto.ExtractTransactionsResponse{
		Report: *recapReport,
	}, nil
}
