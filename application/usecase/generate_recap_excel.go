package usecase

import (
	"context"
	"fmt"

	"sandbox/application/dto"
	"sandbox/infrastructure/excel"
)

// GenerateRecapExcelUseCase handles the business use case for generating the Excel recap
type GenerateRecapExcelUseCase struct {
	excelGenerator *excel.Generator
}

// NewGenerateRecapExcelUseCase creates a new use case instance
func NewGenerateRecapExcelUseCase(excelGenerator *excel.Generator) *GenerateRecapExcelUseCase {
	return &GenerateRecapExcelUseCase{
		excelGenerator: excelGenerator,
	}
}

// Execute performs the Excel recap generation use case
func (uc *GenerateRecapExcelUseCase) Execute(ctx context.Context, req dto.GenerateRecapExcelRequest) (*dto.GenerateRecapExcelResponse, error) {
	// Generate the Excel file
	excelBuffer, err := uc.excelGenerator.GenerateRecapExcel(req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate excel file: %w", err)
	}

	return &dto.GenerateRecapExcelResponse{
		FileContent: excelBuffer.Bytes(),
	}, nil
}
