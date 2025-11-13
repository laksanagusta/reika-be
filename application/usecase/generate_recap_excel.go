package usecase

import (
	"context"
	"fmt"

	"sandbox/application/dto"
	"sandbox/infrastructure/excel"
)

type GenerateRecapExcelUseCase struct {
	excelGenerator *excel.Generator
}

func NewGenerateRecapExcelUseCase(excelGenerator *excel.Generator) *GenerateRecapExcelUseCase {
	return &GenerateRecapExcelUseCase{
		excelGenerator: excelGenerator,
	}
}

func (uc *GenerateRecapExcelUseCase) Execute(ctx context.Context, req dto.RecapReportDTO) (*dto.GenerateRecapExcelResponse, error) {
	excelBuffer, err := uc.excelGenerator.GenerateRecapExcel(req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate excel file: %w", err)
	}

	return &dto.GenerateRecapExcelResponse{
		FileContent: excelBuffer.Bytes(),
	}, nil
}
