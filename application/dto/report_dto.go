package dto

import (
	"time"
)

// ExcelReportRow represents a single row in the Excel report
type ExcelReportRow struct {
	No      int
	Name    string
	NIP     string
	Jabatan string
	Gol     string
	Tujuan  string
	Tanggal time.Time

	// Uang Harian
	UangHarianJmlHari int32
	UangHarianPerHari int32
	UangHarianJumlah  int32

	// Penginapan
	PenginapanJmlMalam int32
	PenginapanPerMalam int32
	PenginapanJumlah   int32

	// Transport
	TransportTiketPesawat int32
	TransportAsal         string
	TransportDaerah       string
	TransportDarat        string
	TransportJumlah       int32

	JumlahDibayarkan int32
}

// GenerateReportRequest represents the request to generate an Excel report.
// It will receive the list of transactions from the frontend.
type GenerateReportRequest struct {
	Transactions []TransactionDTO `json:"transactions"`
}

// GenerateReportResponse represents the response for generating an Excel report.
type GenerateReportResponse struct {
	FileName    string `json:"file_name"`
	FileContent []byte `json:"file_content"`
}

// ReportService represents an interface for report generation related services
type ReportService interface {
	AggregateTransactions(transactions []TransactionDTO) ([]ExcelReportRow, error)
	GenerateExcel(rows []ExcelReportRow) ([]byte, error)
}
