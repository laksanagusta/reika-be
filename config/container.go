package config

import (
	"sandbox/application/usecase"
	"sandbox/domain/transaction"
	"sandbox/infrastructure/excel"
	"sandbox/infrastructure/file"
	"sandbox/infrastructure/gemini"
	"sandbox/interfaces/http/handler"
)

// Container holds all application dependencies
type Container struct {
	// Handlers
	TransactionHandler *handler.TransactionHandler

	// Use Cases
	ExtractTransactionsUseCase *usecase.ExtractTransactionsUseCase
	GenerateRecapExcelUseCase  *usecase.GenerateRecapExcelUseCase

	// Services
	TransactionService *transaction.Service

	// Repositories
	GeminiClient *gemini.Client

	// Processors
	FileProcessor  *file.Processor
	ExcelGenerator *excel.Generator
}

// NewContainer creates and wires up all dependencies
func NewContainer(cfg *Config) *Container {
	// Infrastructure layer
	geminiClient := gemini.NewClient(cfg.Gemini.APIKey)
	fileProcessor := file.NewProcessor()
	excelGenerator := excel.NewGenerator()

	// Domain layer
	transactionService := transaction.NewService(geminiClient)

	// Application layer
	extractTransactionsUseCase := usecase.NewExtractTransactionsUseCase(transactionService)
	generateRecapExcelUseCase := usecase.NewGenerateRecapExcelUseCase(excelGenerator)

	// Interface layer
	transactionHandler := handler.NewTransactionHandler(extractTransactionsUseCase, fileProcessor, generateRecapExcelUseCase)

	return &Container{
		TransactionHandler:         transactionHandler,
		ExtractTransactionsUseCase: extractTransactionsUseCase,
		GenerateRecapExcelUseCase:  generateRecapExcelUseCase,
		TransactionService:         transactionService,
		GeminiClient:               geminiClient,
		FileProcessor:              fileProcessor,
		ExcelGenerator:             excelGenerator,
	}
}
