package gemini

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"sandbox/application/dto"
	"sandbox/domain/transaction"
)

const (
	geminiAPIURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent"
)

// Client represents a Gemini API client
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new Gemini API client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 300 * time.Second, // Increased to 5 minutes for large document processing
		},
	}
}

// ExtractFromDocuments implements the ExtractorRepository interface
func (c *Client) ExtractFromDocuments(ctx context.Context, documents []transaction.Document) (*dto.RecapReportDTO, error) {
	if len(documents) == 0 {
		return nil, errors.New("no documents provided")
	}

	// Check if API key is configured
	if c.apiKey == "" {
		return nil, errors.New("GEMINI_API_KEY is not configured. Please set the environment variable and restart the application")
	}

	// Check if context is already cancelled
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("context cancelled before starting API call: %w", err)
	}

	prompt := c.buildPrompt()
	parts := []map[string]interface{}{
		{"text": prompt},
	}

	// Add all documents to the request
	for _, doc := range documents {
		base64Content := base64.StdEncoding.EncodeToString(doc.Content)
		parts = append(parts, map[string]interface{}{
			"inline_data": map[string]interface{}{
				"mime_type": doc.MimeType,
				"data":      base64Content,
			},
		})
	}

	body := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": parts,
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.getAPIURL(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Provide more detailed error information
		if ctx.Err() != nil {
			return nil, fmt.Errorf("context cancelled during API call: %w", ctx.Err())
		}
		return nil, fmt.Errorf("failed to call Gemini API: %w", err)
	}
	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gemini api error (status %d): %s", resp.StatusCode, string(bodyResp))
	}

	return c.parseResponse(bodyResp)
}

func (c *Client) getAPIURL() string {
	return fmt.Sprintf("%s?key=%s", geminiAPIURL, c.apiKey)
}

func (c *Client) buildPrompt() string {
	return `Baca semua dokumen berikut (gambar atau PDF).
Ekstrak setiap transaksi dan tampilkan dalam format JSON valid berikut ini:

{
  "startDate": "YYYY-MM-DD", -> ambil dari file surat tugas
  "endDate": "YYYY-MM-DD", -> ambil dari file surat tugas
  "activityPurpose": "TUJUAN_AKTIVITAS", -> ambil dari file surat tugas
  "destinationCity": "KOTA_TUJUAN", -> ambil dari file surat tugas
  "spdDate": "YYYY-MM-DD", -> ambil dari file surat tugas
  "departureDate": "YYYY-MM-DD", -> ambil dari file surat tugas
  "returnDate": "YYYY-MM-DD", -> ambil dari file surat tugas
  "receiptSignatureDate": "Tanggal hari ini atau tanggal yang paling baru atau tanggal berjalan",
  "assignees": [
    {
      "name": "NAMA_PEGAWAI", -> ambil dari file surat tugas
      "spd_number": "NOMOR_SPD", -> ambil dari file surat tugas
      "employee_id": "NIP_PEGAWAI", -> ambil dari file surat tugas
      "position": "JABATAN_PEGAWAI", -> ambil dari file surat tugas
      "rank": "GOLONGAN_PEGAWAI", -> ambil dari file surat tugas
      "transactions": [
        {
          "name": "NAMA_PEMESAN_TRANSAKSI",
          "type": "accommodation | transport | other",
          "subtype": "hotel | flight | train | taxi | ...",
          "amount": number,
          "total_night": number,
          "subtotal": number, -> hasil amount*total_night kalo dia accomodation tapi kalo selain itu langsung ambil dari amount aja
	      "description" : string, -> ini adalah keterangan transaksi ini transaksi apa, misalkan gojek dari alamat1 ke alamat2, kalo hotel jelasin juga hotelnya
	      "transport_detail" : string, -> ini terisi hanya jika dia transport darat ya (pesawat tidak termasuk) 1.jika dia dari bandara soetta atau tujuannya ke bandara soetta maka valuenya menjadi "transport_asal" atau kalau dia transportasinya di jakarta juga masuk trasnport asal 2.jika mengandung bandara lain selain soetta maka valuenya adalah "transport_daerah"
        }
      ]
    }
  ]
}

- Kembalikan hasil hanya dalam JSON valid (tanpa teks tambahan).
- Jangan bungkus JSON dengan tanda kutip atau karakter escape.
- Jika total_night tidak ada, field tersebut boleh dihapus.
- Pastikan angka hanya berupa digit (tanpa simbol mata uang).
- Untuk data transaksi, nama yang digunakan harus sesuai dengan nama yang tercantum di surat tugas. Harap lakukan pengecekan dan pencocokan dengan surat tugas.
- Jika nama pemesan di transaksi tersebut tidak tercantum di surat tugas, mohon assign ke salah satu nama yang ada di surat tugas.
- Jangan menggunakan nama driver sebagai nama transaksi — gunakan nama pemesan.
- Group semua transaksi di bawah setiap assignee.
`
}

func (c *Client) parseResponse(bodyResp []byte) (*dto.RecapReportDTO, error) {
	var geminiAPIResponse geminiResponse
	if err := json.Unmarshal(bodyResp, &geminiAPIResponse); err != nil {
		return nil, fmt.Errorf("failed to parse Gemini API response wrapper: %w", err)
	}

	if len(geminiAPIResponse.Candidates) == 0 || len(geminiAPIResponse.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New("empty response candidates or parts from Gemini API")
	}

	rawText := geminiAPIResponse.Candidates[0].Content.Parts[0].Text
	cleanJSON := c.cleanJSON(rawText)

	var geminiRawReport geminiReportResponse
	if err := json.Unmarshal([]byte(cleanJSON), &geminiRawReport); err != nil {
		return nil, fmt.Errorf("failed to parse Gemini report content: %w (raw: %s)", err, cleanJSON)
	}

	// Convert raw assignee responses to dto.AssigneeDTO
	assignees := make([]dto.AssigneeDTO, 0, len(geminiRawReport.Assignees))
	for _, rawAssignee := range geminiRawReport.Assignees {
		transactionsDTO := make([]dto.TransactionDTO, 0, len(rawAssignee.Transactions))
		for _, rawTx := range rawAssignee.Transactions {
			transactionsDTO = append(transactionsDTO, dto.TransactionDTO{
				Name:            rawTx.Name,
				Type:            rawTx.Type,
				Subtype:         rawTx.Subtype,
				Amount:          rawTx.Amount,
				TotalNight:      rawTx.TotalNight,
				Subtotal:        rawTx.Subtotal,
				PaymentType:     "", // Assuming default empty, needs to be derived if applicable
				Description:     rawTx.Description,
				TransportDetail: rawTx.TransportDetail,
			})
		}

		assignees = append(assignees, dto.AssigneeDTO{
			Name:         rawAssignee.Name,
			SpdNumber:    rawAssignee.SpdNumber,
			EmployeeID:   rawAssignee.EmployeeID,
			Position:     rawAssignee.Position,
			Rank:         rawAssignee.Rank,
			Transactions: transactionsDTO,
		})
	}

	return &dto.RecapReportDTO{
		StartDate:            geminiRawReport.StartDate,
		EndDate:              geminiRawReport.EndDate,
		ActivityPurpose:      geminiRawReport.ActivityPurpose,
		DestinationCity:      geminiRawReport.DestinationCity,
		SpdDate:              geminiRawReport.SpdDate,
		DepartureDate:        geminiRawReport.DepartureDate,
		ReturnDate:           geminiRawReport.ReturnDate,
		ReceiptSignatureDate: geminiRawReport.ReceiptSignatureDate,
		Assignees:            assignees,
	}, nil
}

func (c *Client) cleanJSON(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "```json")
	s = strings.TrimPrefix(s, "```JSON")
	s = strings.TrimSuffix(s, "```")
	return strings.TrimSpace(s)
}

// Internal types for JSON parsing
type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// geminiReportResponse represents the full report structure from Gemini API
type geminiReportResponse struct {
	StartDate            string                `json:"startDate"`
	EndDate              string                `json:"endDate"`
	ActivityPurpose      string                `json:"activityPurpose"`
	DestinationCity      string                `json:"destinationCity"`
	SpdDate              string                `json:"spdDate"`
	DepartureDate        string                `json:"departureDate"`
	ReturnDate           string                `json:"returnDate"`
	ReceiptSignatureDate string                `json:"receiptSignatureDate"`
	Assignees            []rawAssigneeResponse `json:"assignees"`
}

type rawAssigneeResponse struct {
	Name         string           `json:"name"`
	SpdNumber    string           `json:"spd_number"`
	EmployeeID   string           `json:"employee_id"`
	Position     string           `json:"position"`
	Rank         string           `json:"rank"`
	Transactions []rawTransaction `json:"transactions"`
}

type rawTransaction struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Subtype         string `json:"subtype"`
	Amount          int32  `json:"amount"`
	TotalNight      *int32 `json:"total_night,omitempty"`
	Subtotal        int32  `json:"subtotal"`
	Description     string `json:"description"`
	TransportDetail string `json:"transport_detail"`
}
