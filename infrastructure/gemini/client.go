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
			Timeout: 60 * time.Second,
		},
	}
}

// ExtractFromDocuments implements the ExtractorRepository interface
func (c *Client) ExtractFromDocuments(ctx context.Context, documents []transaction.Document) ([]*transaction.Transaction, error) {
	if len(documents) == 0 {
		return nil, errors.New("no documents provided")
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
Ekstrak setiap transaksi dan tampilkan dalam format ARRAY JSON valid berikut ini:

[
  {
    "name": "NAMA_ORANG",
    "type": "accommodation | transport | other",
    "subtype": "hotel | flight | train | taxi | ...",
    "amount": number,
    "total_night": number,
    "subtotal": number, -> hasil amount*total_night kalo dia accomodation tapi kalo selain itu langsung ambil dari amount aja
	"description" : string, -> ini adalah keterangan transaksi ini transaksi apa, misalkan gojek dari alamat1 ke alamat2, kalo hotel jelasin juga hotelnya
	"transport_detail" : string, -> ini terisi hanya jika dia transport darat ya (pesawat tidak termasuk) 1.jika dia dari bandara soetta atau tujuannya ke bandara soetta maka valuenya menjadi "transport_asal" 2.jika mengandung bandara lain selain soetta maka valuenya adalah "transport_daerah"
	"employee_id" : string, -> ini adalah NIP ambil dari surat tugas
	"position" : string, -> ini adalah jabatan ambil dari surat tugas
	"rank" : string -> ini adalah golongan ambil dari surat tugas
  }
]

- Kembalikan hasil hanya dalam JSON array valid (tanpa teks tambahan).
- Jangan bungkus JSON dengan tanda kutip atau karakter escape.
- Jika total_night tidak ada, field tersebut boleh dihapus.
- Pastikan angka hanya berupa digit (tanpa simbol mata uang).
- Gabungkan semua nota dari semua file dalam satu array.
- Nama orangnya cukup nama asli aja ya tanpa ada embel-embel lainnya
- untuk data transaksinya harusnya nama yg muncul adalah nama yg ada di surat tugas jadi harap cocokkan sesuai surat tugas
`
}

func (c *Client) parseResponse(bodyResp []byte) ([]*transaction.Transaction, error) {
	var geminiResp geminiResponse
	if err := json.Unmarshal(bodyResp, &geminiResp); err != nil {
		return nil, fmt.Errorf("failed to parse Gemini response: %w", err)
	}

	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New("empty response from Gemini")
	}

	rawText := geminiResp.Candidates[0].Content.Parts[0].Text
	cleanJSON := c.cleanJSON(rawText)

	var rawTransactions []rawTransaction
	if err := json.Unmarshal([]byte(cleanJSON), &rawTransactions); err != nil {
		return nil, fmt.Errorf("failed to parse transactions JSON: %w (raw: %s)", err, cleanJSON)
	}

	// Convert raw transactions to domain entities
	transactions := make([]*transaction.Transaction, 0, len(rawTransactions))
	for _, raw := range rawTransactions {
		tx, err := transaction.NewTransaction(
			raw.Name,
			raw.Type,
			raw.Subtype,
			raw.Amount,
			raw.Subtotal,
			raw.TotalNight,
			raw.Description,
			raw.TransportDetail,
			raw.EmployeeID,
			raw.Position,
			raw.Rank,
		)
		if err != nil {
			// Log error but continue processing other transactions
			continue
		}
		transactions = append(transactions, tx)
	}

	return transactions, nil
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

type rawTransaction struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Subtype         string `json:"subtype"`
	Amount          int32  `json:"amount"`
	TotalNight      *int32 `json:"total_night,omitempty"`
	Subtotal        int32  `json:"subtotal"`
	Description     string `json:"description"`
	TransportDetail string `json:"transport_detail"`
	EmployeeID      string `json:"employee_id"`
	Position        string `json:"position"`
	Rank            string `json:"rank"`
}
