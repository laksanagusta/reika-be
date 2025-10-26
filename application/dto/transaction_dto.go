package dto

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/invopop/validation"
)

var dateFormatRegex = regexp.MustCompile(`^\d{1,2}\s+(Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember)\s+\d{4}$`)

// parseIndonesianDate parses Indonesian date format (e.g., "25 Oktober 2025")
func parseIndonesianDate(dateStr string) (time.Time, error) {
	// Extract day, month, and year
	parts := strings.Fields(dateStr)
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid date format")
	}

	dayStr := parts[0]
	monthStr := parts[1]
	yearStr := parts[2]

	// Parse day
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid day: %s", dayStr)
	}

	// Parse month
	month, exists := IndonesianMonths[monthStr]
	if !exists {
		return time.Time{}, fmt.Errorf("invalid month: %s", monthStr)
	}

	// Parse year
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid year: %s", yearStr)
	}

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), nil
}

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

func (tx *TransactionDTO) Validate(fieldPrefix string) error {
	// Basic validation
	if err := validation.ValidateStruct(tx,
		validation.Field(&tx.Type, validation.Required),
		validation.Field(&tx.Subtype, validation.Length(0, 100)),
		validation.Field(&tx.Amount, validation.Required, validation.Min(0)),
		validation.Field(&tx.Subtotal, validation.Required, validation.Min(0)),
		validation.Field(&tx.PaymentType, validation.Length(0, 50)),
		validation.Field(&tx.Description, validation.Length(0, 500)),
		validation.Field(&tx.TransportDetail, validation.Length(0, 200)),
	); err != nil {
		return err
	}

	// Validate transaction type
	if !tx.isValidTransactionType() {
		return validation.NewError(fieldPrefix+".type", "type must be one of: accommodation, transport, other")
	}

	// Conditional validation for accommodation
	if strings.ToLower(tx.Type) == "accommodation" && tx.TotalNight != nil && *tx.TotalNight <= 0 {
		return validation.NewError(fieldPrefix+".total_night", "total_night must be positive for accommodation transactions")
	}

	return nil
}

func (tx *TransactionDTO) isValidTransactionType() bool {
	normalizedType := strings.ToLower(tx.Type)
	switch normalizedType {
	case "accommodation", "transport", "other":
		return true
	}
	return false
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

func (a *AssigneeDTO) Validate(index int) error {
	fieldPrefix := fmt.Sprintf("assignees[%d]", index)

	if err := validation.ValidateStruct(a,
		validation.Field(&a.Name, validation.Required, validation.Length(1, 100)),
		validation.Field(&a.SpdNumber, validation.Required, validation.Length(1, 50)),
		validation.Field(&a.EmployeeID, validation.Required, validation.Length(1, 50)),
		validation.Field(&a.Position, validation.Required, validation.Length(1, 100)),
		validation.Field(&a.Rank, validation.Required, validation.Length(1, 50)),
		validation.Field(&a.Transactions, validation.Required),
	); err != nil {
		return err
	}

	// Validate transactions
	if len(a.Transactions) == 0 {
		return validation.NewError(fieldPrefix+".transactions", "at least one transaction is required for each assignee")
	}

	for i, tx := range a.Transactions {
		if err := tx.Validate(fmt.Sprintf("%s.transactions[%d]", fieldPrefix, i)); err != nil {
			return err
		}
	}

	return nil
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

func (r *RecapReportDTO) Validate() error {
	// Validate main report structure
	if err := validation.ValidateStruct(r,
		validation.Field(&r.StartDate, validation.Required, validation.Match(dateFormatRegex)),
		validation.Field(&r.EndDate, validation.Required, validation.Match(dateFormatRegex)),
		validation.Field(&r.ActivityPurpose, validation.Required, validation.Length(1, 200)),
		validation.Field(&r.DestinationCity, validation.Required, validation.Length(1, 100)),
		validation.Field(&r.SpdDate, validation.Required, validation.Match(dateFormatRegex)),
		validation.Field(&r.DepartureDate, validation.Required, validation.Match(dateFormatRegex)),
		validation.Field(&r.ReturnDate, validation.Required, validation.Match(dateFormatRegex)),
		validation.Field(&r.ReceiptSignatureDate, validation.Match(dateFormatRegex)),
		validation.Field(&r.Assignees, validation.Required),
	); err != nil {
		return err
	}

	// Validate date logic
	if err := r.validateDateLogic(); err != nil {
		return err
	}

	// Validate assignees
	if len(r.Assignees) == 0 {
		return validation.NewError("assignees", "at least one assignee is required")
	}

	for i, assignee := range r.Assignees {
		if err := assignee.Validate(i); err != nil {
			return err
		}
	}

	return nil
}

func (r *RecapReportDTO) validateDateLogic() error {
	startDate, err := parseIndonesianDate(r.StartDate)
	if err != nil {
		return validation.NewError("startDate", fmt.Sprintf("invalid start date format: %v (expected format: '25 Oktober 2025')", err))
	}

	endDate, err := parseIndonesianDate(r.EndDate)
	if err != nil {
		return validation.NewError("endDate", fmt.Sprintf("invalid end date format: %v (expected format: '25 Oktober 2025')", err))
	}

	_, err = parseIndonesianDate(r.SpdDate)
	if err != nil {
		return validation.NewError("spdDate", fmt.Sprintf("invalid SPD date format: %v (expected format: '25 Oktober 2025')", err))
	}

	departureDate, err := parseIndonesianDate(r.DepartureDate)
	if err != nil {
		return validation.NewError("departureDate", fmt.Sprintf("invalid departure date format: %v (expected format: '25 Oktober 2025')", err))
	}

	returnDate, err := parseIndonesianDate(r.ReturnDate)
	if err != nil {
		return validation.NewError("returnDate", fmt.Sprintf("invalid return date format: %v (expected format: '25 Oktober 2025')", err))
	}

	// Validate date relationships
	if startDate.After(endDate) {
		return validation.NewError("date_range", "start date must be before or equal to end date")
	}

	if departureDate.After(returnDate) {
		return validation.NewError("travel_dates", "departure date must be before or equal to return date")
	}

	// Receipt signature date validation (if provided)
	if r.ReceiptSignatureDate != "" {
		receiptDate, err := parseIndonesianDate(r.ReceiptSignatureDate)
		if err != nil {
			return validation.NewError("receiptSignatureDate", fmt.Sprintf("invalid receipt signature date format: %v (expected format: '25 Oktober 2025')", err))
		}
		if receiptDate.Before(returnDate) {
			return validation.NewError("receiptSignatureDate", "receipt signature date must be after or equal to return date")
		}
	}

	return nil
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
