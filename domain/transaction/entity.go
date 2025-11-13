package transaction

import (
	"strings"
)

type TransactionType string

const (
	TransactionTypeAccommodation TransactionType = "accommodation"
	TransactionTypeTransport     TransactionType = "transport"
	TransactionTypeOther         TransactionType = "other"
)

type Transaction struct {
	Name            string
	TxType          TransactionType
	Subtype         string
	Amount          int32
	TotalNight      *int32
	Subtotal        int32
	Description     string
	TransportDetail string
	EmployeeID      string
	Position        string
	Rank            string
}

func NewTransaction(name, txType, subtype string, amount, subtotal int32, totalNight *int32, description string, transportDetail string, employeeID, position, rank string) (*Transaction, error) {
	validType := TransactionType(strings.ToLower(txType))
	if !isValidTransactionType(validType) {
		validType = TransactionTypeOther
	}

	return &Transaction{
		Name:            strings.TrimSpace(name),
		TxType:          validType,
		Subtype:         strings.TrimSpace(subtype),
		Amount:          amount,
		TotalNight:      totalNight,
		Subtotal:        subtotal,
		Description:     description,
		TransportDetail: transportDetail,
		EmployeeID:      employeeID,
		Position:        position,
		Rank:            rank,
	}, nil
}

// Getters
func (t *Transaction) GetName() string {
	return t.Name
}

func (t *Transaction) GetType() TransactionType {
	return t.TxType
}

func (t *Transaction) GetSubtype() string {
	return t.Subtype
}

func (t *Transaction) GetAmount() int32 {
	return t.Amount
}

func (t *Transaction) GetTotalNight() *int32 {
	return t.TotalNight
}

func (t *Transaction) GetSubtotal() int32 {
	return t.Subtotal
}

func (t *Transaction) GetTransportDetail() string {
	return t.TransportDetail
}

func (t *Transaction) GetDescription() string {
	return t.Description
}

func (t *Transaction) GetEmployeeID() string {
	return t.EmployeeID
}

func (t *Transaction) GetPosition() string {
	return t.Position
}

func (t *Transaction) GetRank() string {
	return t.Rank
}

func (t *Transaction) IsAccommodation() bool {
	return t.TxType == TransactionTypeAccommodation
}

func (t *Transaction) IsTransport() bool {
	return t.TxType == TransactionTypeTransport
}

func (t *Transaction) CalculateTotal() int32 {
	if t.TotalNight != nil && *t.TotalNight > 0 {
		return t.Amount * *t.TotalNight
	}
	return t.Subtotal
}

func isValidTransactionType(t TransactionType) bool {
	switch t {
	case TransactionTypeAccommodation, TransactionTypeTransport, TransactionTypeOther:
		return true
	}
	return false
}
