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
	SpdNumber       string `json:"spd_number"`
	Description     string `json:"description"`
	TransportDetail string `json:"transport_detail"`
	EmployeeID      string `json:"employee_id"`
	Position        string `json:"position"`
	Rank            string `json:"rank"`
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
	Transactions []TransactionDTO `json:"transactions"`
	Count        int              `json:"count"`
}

// GenerateRecapExcelRequest represents the request for generating the recap Excel
type GenerateRecapExcelRequest struct {
	StartDate       string           `json:"start_date"`
	EndDate         string           `json:"end_date"`
	DepartureDate   string           `json:"departure_date"`
	ReceiptSignDate string           `json:"receipt_sign_date"`
	ReturnDate      string           `json:"return_date"`
	SpdDate         string           `json:"spd_date"`
	Destination     string           `json:"destination"`
	DestinationCity string           `json:"destination_city"`
	Transactions    []TransactionDTO `json:"transactions"`
}

// GenerateRecapExcelResponse represents the response for generating the recap Excel
type GenerateRecapExcelResponse struct {
	FileContent []byte `json:"file_content"`
}
