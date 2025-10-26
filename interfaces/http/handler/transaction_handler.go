package handler

import (
	"log"

	"sandbox/application/dto"
	"sandbox/application/usecase"
	"sandbox/infrastructure/file"

	"github.com/gofiber/fiber/v2"
)

// TransactionHandler handles HTTP requests for transactions
type TransactionHandler struct {
	extractUseCase            *usecase.ExtractTransactionsUseCase
	fileProcessor             *file.Processor
	generateRecapExcelUseCase *usecase.GenerateRecapExcelUseCase
}

// NewTransactionHandler creates a new transaction handler
func NewTransactionHandler(
	extractUseCase *usecase.ExtractTransactionsUseCase,
	fileProcessor *file.Processor,
	generateRecapExcelUseCase *usecase.GenerateRecapExcelUseCase,
) *TransactionHandler {
	return &TransactionHandler{
		extractUseCase:            extractUseCase,
		fileProcessor:             fileProcessor,
		generateRecapExcelUseCase: generateRecapExcelUseCase,
	}
}

// UploadAndExtract handles the file upload and extraction endpoint
func (h *TransactionHandler) UploadAndExtract(c *fiber.Ctx) error {
	log.Println("Processing upload request")

	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form data",
		})
	}

	fileHeaders := form.File["file"]
	if len(fileHeaders) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No files uploaded",
		})
	}

	// Process uploaded files
	processedFiles, err := h.fileProcessor.ProcessMultipleFiles(fileHeaders)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Convert to DTO
	fileUploads := make([]dto.FileUpload, len(processedFiles))
	for i, pf := range processedFiles {
		fileUploads[i] = dto.FileUpload{
			Content:  pf.Content,
			Filename: pf.Filename,
			MimeType: pf.MimeType,
		}
	}

	// Execute use case
	request := dto.ExtractTransactionsRequest{
		Files: fileUploads,
	}

	response, err := h.extractUseCase.Execute(c.Context(), request)
	if err != nil {
		log.Printf("Error extracting transactions: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to extract transactions",
			"details": err.Error(),
		})
	}

	// Return transactions only (backward compatible)
	return c.JSON(response.Transactions)
}

func (h *TransactionHandler) GenerateRecapExcel(c *fiber.Ctx) error {
	log.Println("Generating Excel recap file")

	var reqBody dto.GenerateRecapExcelRequest
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	if len(reqBody.Transactions) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No transactions provided for Excel generation",
		})
	}

	response, err := h.generateRecapExcelUseCase.Execute(c.Context(), reqBody)
	if err != nil {
		log.Printf("Error generating Excel recap: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to generate Excel recap file",
			"details": err.Error(),
		})
	}

	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", "attachment; filename=\"kwitansi-perjadin.xlsx\"")
	return c.Send(response.FileContent)
}

// UploadAndExtractDetailed returns detailed response with count
func (h *TransactionHandler) UploadAndExtractDetailed(c *fiber.Ctx) error {
	log.Println("Processing upload request (detailed)")

	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form data",
		})
	}

	fileHeaders := form.File["file"]
	if len(fileHeaders) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No files uploaded",
		})
	}

	// Process uploaded files
	processedFiles, err := h.fileProcessor.ProcessMultipleFiles(fileHeaders)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Convert to DTO
	fileUploads := make([]dto.FileUpload, len(processedFiles))
	for i, pf := range processedFiles {
		fileUploads[i] = dto.FileUpload{
			Content:  pf.Content,
			Filename: pf.Filename,
			MimeType: pf.MimeType,
		}
	}

	// Execute use case
	request := dto.ExtractTransactionsRequest{
		Files: fileUploads,
	}

	response, err := h.extractUseCase.Execute(c.Context(), request)
	if err != nil {
		log.Printf("Error extracting transactions: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to extract transactions",
			"details": err.Error(),
		})
	}

	// Return full response
	return c.JSON(response)
}
