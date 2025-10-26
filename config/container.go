package config

import (
	"sandbox/application/usecase"
	domainMeeting "sandbox/domain/meeting"
	"sandbox/domain/transaction"
	"sandbox/infrastructure/drive"
	"sandbox/infrastructure/excel"
	"sandbox/infrastructure/file"
	"sandbox/infrastructure/gemini"
	meetingInfra "sandbox/infrastructure/meeting"
	"sandbox/infrastructure/notification"
	"sandbox/infrastructure/zoom"
	"sandbox/interfaces/http/handler"
)

// Container holds all application dependencies
type Container struct {
	// Handlers
	TransactionHandler *handler.TransactionHandler
	MeetingHandler     *handler.MeetingHandler

	// Use Cases
	ExtractTransactionsUseCase *usecase.ExtractTransactionsUseCase
	GenerateRecapExcelUseCase  *usecase.GenerateRecapExcelUseCase
	CreateMeetingUseCase       *usecase.CreateMeetingUseCase

	// Services
	TransactionService *transaction.Service
	MeetingService     *domainMeeting.Service

	// Repositories
	GeminiClient    *gemini.Client
	MeetingRepo     domainMeeting.Repository

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

	// Meeting infrastructure
	zoomClient := zoom.NewClient(cfg.Zoom.APIKey, cfg.Zoom.APISecret)
	driveClient := drive.NewClient(cfg.Drive.APIKey)
	notificationClient := notification.NewClient(cfg.Notification.APIKey)
	meetingRepo := meetingInfra.NewRepository(zoomClient, driveClient, notificationClient)

	// Domain layer
	transactionService := transaction.NewService(geminiClient)
	meetingService := domainMeeting.NewService(meetingRepo)

	// Application layer
	extractTransactionsUseCase := usecase.NewExtractTransactionsUseCase(transactionService)
	generateRecapExcelUseCase := usecase.NewGenerateRecapExcelUseCase(excelGenerator)
	createMeetingUseCase := usecase.NewCreateMeetingUseCase(meetingService)

	// Interface layer
	transactionHandler := handler.NewTransactionHandler(extractTransactionsUseCase, fileProcessor, generateRecapExcelUseCase)
	meetingHandler := handler.NewMeetingHandler(createMeetingUseCase)

	return &Container{
		TransactionHandler:         transactionHandler,
		MeetingHandler:             meetingHandler,
		ExtractTransactionsUseCase: extractTransactionsUseCase,
		GenerateRecapExcelUseCase:  generateRecapExcelUseCase,
		CreateMeetingUseCase:       createMeetingUseCase,
		TransactionService:         transactionService,
		MeetingService:             meetingService,
		GeminiClient:               geminiClient,
		MeetingRepo:                meetingRepo,
		FileProcessor:              fileProcessor,
		ExcelGenerator:             excelGenerator,
	}
}
