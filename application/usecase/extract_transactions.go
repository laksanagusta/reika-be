package usecase

import (
	"context"

	"sandbox/application/dto"
	"sandbox/domain/transaction"
)

type ExtractTransactionsUseCase struct {
	transactionService *transaction.Service
}

func NewExtractTransactionsUseCase(transactionService *transaction.Service) *ExtractTransactionsUseCase {
	return &ExtractTransactionsUseCase{
		transactionService: transactionService,
	}
}

func (uc *ExtractTransactionsUseCase) Execute(ctx context.Context, req dto.ExtractTransactionsRequest) (*dto.ExtractTransactionsResponse, error) {
	documents := make([]transaction.Document, len(req.Files))
	for i, file := range req.Files {
		documents[i] = transaction.Document{
			Content:  file.Content,
			MimeType: file.MimeType,
			Filename: file.Filename,
		}
	}

	recapReport, err := uc.transactionService.ExtractTransactions(ctx, documents)
	if err != nil {
		return nil, err
	}

	return &dto.ExtractTransactionsResponse{
		Report: *recapReport,
	}, nil
}
