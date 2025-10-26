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
	transactions, err := uc.transactionService.ExtractTransactions(ctx, documents)
	if err != nil {
		return nil, err
	}

	// Convert domain objects to DTOs
	transactionDTOs := make([]dto.TransactionDTO, len(transactions))
	for i, tx := range transactions {
		transactionDTOs[i] = dto.TransactionDTO{
			Name:            tx.GetName(),
			Type:            string(tx.GetType()),
			Subtype:         tx.GetSubtype(),
			Amount:          tx.GetAmount(),
			TotalNight:      tx.GetTotalNight(),
			Subtotal:        tx.GetSubtotal(),
			Description:     tx.GetDescription(),
			TransportDetail: tx.GetTransportDetail(),
			EmployeeID:      tx.GetEmployeeID(),
			Position:        tx.GetPosition(),
			Rank:            tx.GetRank(),
		}
	}

	return &dto.ExtractTransactionsResponse{
		Transactions: transactionDTOs,
		Count:        len(transactionDTOs),
	}, nil
}
