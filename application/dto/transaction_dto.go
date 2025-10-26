package dto

// TransactionDTO represents the data transfer object for transactions
type TransactionDTO struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Subtype         string `json:"subtype"`
	Amount          int32  `json:"amount"`
	TotalNight      *int32 `json:"total_night,omitempty"`
	Subtotal        int32  `json:"subtotal"`
	PaymentType     string `json:"payment_type"`
	Description     string `json:"description"`
	TransportDetail string `json:"transport_detail"`
}

// AssigneeDTO represents an assignee with their transactions
type AssigneeDTO struct {
	Name         string           `json:"name"`
	SpdNumber    string           `json:"spd_number"`
	EmployeeID   string           `json:"employee_id"`
	Position     string           `json:"position"`
	Rank         string           `json:"rank"`
	Transactions []TransactionDTO `json:"transactions"`
}

// RecapReportDTO represents the overall structure of the recap report
type RecapReportDTO struct {
	StartDate            string        `json:"startDate"`
	EndDate              string        `json:"endDate"`
	ActivityPurpose      string        `json:"activityPurpose"` // This maps to Destination in current GenerateRecapExcelRequest
	DestinationCity      string        `json:"destinationCity"`
	SpdDate              string        `json:"spdDate"`
	DepartureDate        string        `json:"departureDate"`
	ReturnDate           string        `json:"returnDate"`
	ReceiptSignatureDate string        `json:"receiptSignatureDate"`
	Assignees            []AssigneeDTO `json:"assignees"`
}

// ExtractTransactionsRequest represents the request for extracting transactions
type ExtractTransactionsRequest struct {
	Files []FileUpload
}

// FileUpload represents an uploaded file
type FileUpload struct {
	Content  []byte
	Filename string
	MimeType string
}

// ExtractTransactionsResponse represents the response
type ExtractTransactionsResponse struct {
	Report RecapReportDTO `json:"report"`
}

// GenerateRecapExcelRequest represents the request for generating the recap Excel
type GenerateRecapExcelRequest struct {
	StartDate            string        `json:"startDate"`
	EndDate              string        `json:"endDate"`
	ActivityPurpose      string        `json:"activityPurpose"`
	DestinationCity      string        `json:"destinationCity"`
	SpdDate              string        `json:"spdDate"`
	DepartureDate        string        `json:"departureDate"`
	ReturnDate           string        `json:"returnDate"`
	ReceiptSignatureDate string        `json:"receiptSignatureDate"`
	Assignees            []AssigneeDTO `json:"assignees"`
}

// GenerateRecapExcelResponse represents the response for generating the recap Excel
type GenerateRecapExcelResponse struct {
	FileContent []byte `json:"file_content"`
}
