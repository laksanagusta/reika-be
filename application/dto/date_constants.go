package dto

import (
	"time"
)

// Indonesian month name to month number mapping
var IndonesianMonths = map[string]time.Month{
	"Januari":   time.January,
	"Februari":  time.February,
	"Maret":     time.March,
	"April":     time.April,
	"Mei":       time.May,
	"Juni":      time.June,
	"Juli":      time.July,
	"Agustus":   time.August,
	"September": time.September,
	"Oktober":   time.October,
	"November":  time.November,
	"Desember":  time.December,
}