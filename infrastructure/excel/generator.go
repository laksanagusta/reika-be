package excel

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"sandbox/application/dto"
	"sandbox/domain/transaction"

	"github.com/xuri/excelize/v2"
)

const (
	workUnit            = "Setditjen Penanggulangan Penyakit"
	commitmentOfficer   = "Ruly Wahyuni, SE, MKM"
	commitmentOfficerId = "197508142000032001"

	expenditureTreasurer   = "Fatmawati Husain, SE, M.Ak"
	expenditureTreasurerId = "198608202005012002"

	payer   = "Dian Pratiwi Andini"
	payerId = "198510312008012003"
)

// Generator handles the generation of Excel files
type Generator struct {
	file   *excelize.File
	styles map[string]int
}

// NewGenerator creates a new Excel generator
func NewGenerator() *Generator {
	file := excelize.NewFile()
	return &Generator{
		file:   file,
		styles: make(map[string]int),
	}
}

func (g *Generator) generateTitle(sheetName string) error {
	if err := g.file.MergeCell(sheetName, "A2", "U2"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "A3", "U3"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "A4", "U4"); err != nil {
		return err
	}

	titleStyle, err := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Bold:   true,
			Size:   12,
			Family: "Tahoma",
		},
	})
	if err != nil {
		return err
	}
	subTitleStyle, err := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Bold:   true,
			Size:   12,
			Family: "Tahoma",
		},
	})
	if err != nil {
		return err
	}
	accountStyle, err := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Bold:   true,
			Size:   12,
			Family: "Tahoma",
		},
	})
	if err != nil {
		return err
	}

	if sheetName == "PEMANTAUAN REKAP RAMPUNG" {
		g.file.SetCellValue(sheetName, "A2", "Rekapitulasi Biaya Perjalanan Dinas Rampung")
	} else {
		g.file.SetCellValue(sheetName, "A2", "Rekapitulasi Uang Muka Biaya Perjalanan Dinas")
	}

	g.file.SetCellValue(sheetName, "A3", "Rekapitulasi Biaya Perjalanan Dinas dalam Rangka Pemantauan dan Evaluasi Pelaksanaan Program di Daerah")
	g.file.SetCellValue(sheetName, "A4", "AKUN : 4815.EBD.953.501.B.524111")

	if err := g.file.SetCellStyle(sheetName, "A2", "U2", titleStyle); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "A3", "U3", subTitleStyle); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "A4", "U4", accountStyle); err != nil {
		return err
	}
	return nil
}

func (g *Generator) generateTableHeader(sheetName string) error {
	headerStyle, err := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Bold:   true,
			Size:   10,
			Family: "Tahoma",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A8", "No"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "A8", "A10"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "B8", "Nama"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "B8", "B10"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "C8", "Employee ID"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "C8", "C10"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "D8", "Position"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "D8", "D10"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "E8", "Rank"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "E8", "E10"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "F8", "Tujuan"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "F8", "F10"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "G8", "Tanggal"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "G8", "G10"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "H8", "Uang Harian"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "H9", "Konstanta:"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "H9", "J9"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "L8", "Penginapan"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "L9", "Konstanta:"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "L9", "N9"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "P8", "Transport"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "P9", "Konstanta:"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "P9", "R9"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "U8", "Jumlah Dibayarkan (Rp)"); err != nil {
		return err
	}

	if sheetName == "KW RAMPUNG" {
		if err := g.file.SetCellValue(sheetName, "Y8", "Jumlah SPJ Rampung (Rp)"); err != nil {
			return err
		}
		if err := g.file.SetCellValue(sheetName, "Z8", "Jumlah SPJ Uang Muka (Rp)"); err != nil {
			return err
		}
		if err := g.file.SetCellValue(sheetName, "AA8", "Jumlah Dibayarkan (Rp)"); err != nil {
			return err
		}
	}

	if err := g.file.SetCellValue(sheetName, "V8", "No SPD"); err != nil {
		return err
	}

	// Sub-headers for Uang Harian
	if err := g.file.SetCellValue(sheetName, "H10", "Jml Hari"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "H10", "I10"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "J10", "Perhari"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "K10", "Jumlah"); err != nil {
		return err
	}

	// Sub-headers for Penginapan
	if err := g.file.SetCellValue(sheetName, "L10", "Jml Hari"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "L10", "M10"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "N10", "Perhari"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O10", "Jumlah"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "P10", "Tiket Pesawat"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "Q10", "Transport Asal"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "R10", "Transport Daerah"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "S10", "Transport Darat"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "T10", "Jumlah"); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "A8", "A9"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "B8", "B9"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "C8", "C9"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "D8", "D9"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "E8", "E9"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "F8", "F9"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "G8", "G9"); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "H8", "K8"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "L8", "O8"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "P8", "T8"); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "A9", "A9"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "Y8", "Y9"); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "U8", "U10"); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "A8", "V10", headerStyle); err != nil {
		return err
	}
	return nil
}

type PersonRecap struct {
	Name                    string
	NIP                     string
	Jabatan                 string
	Gol                     string
	Tujuan                  string
	Tanggal                 string
	NoSpd                   string
	UMUangHarianJmlHari     int32
	UMUangHarianPerhari     int32
	UMUangHarianJumlah      int32
	UMPenginapanJmlHari     int32
	UMPenginapanPerhari     int32
	UMPenginapanJumlah      int32
	UMTransportTiketPesawat int32
	UMTransportJumlah       int32
	UMTotalDibayarkan       int32

	RUangHarianJmlHari     int32
	RUangHarianPerhari     int32
	RUangHarianJumlah      int32
	RPenginapanJmlHari     int32
	RPenginapanPerhari     int32
	RPenginapanJumlah      int32
	RTransportTiketPesawat int32
	RTransportAsal         int32
	RTransportDaerah       int32
	RTransportDarat        int32
	RTransportJumlah       int32
	RTotalDibayarkan       int32
}

func (g *Generator) generateTableData(sheetName string, assignees []dto.AssigneeDTO, currentRow int) (int, error) {
	dataStyle, err := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical: "center",
		},
		Font: &excelize.Font{
			Size:   10,
			Family: "Tahoma",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		return currentRow, err
	}

	personData := make(map[string]*PersonRecap)
	constUangHarianJmlHari := int32(2)
	constUangHarianPerhari := int32(688000)
	constUangHarianJumlah := constUangHarianJmlHari * constUangHarianPerhari

	for _, assignee := range assignees {
		data, exists := personData[assignee.EmployeeID]
		if !exists {
			data = &PersonRecap{
				Name:                assignee.Name,
				NIP:                 assignee.EmployeeID,
				Jabatan:             assignee.Position,
				Gol:                 assignee.Rank,
				Tujuan:              "-", // To be filled from report DTO or transaction description
				Tanggal:             "-", // To be filled from report DTO or transaction date
				NoSpd:               assignee.SpdNumber,
				UMUangHarianJmlHari: constUangHarianJmlHari,
				UMUangHarianPerhari: constUangHarianPerhari,
				UMUangHarianJumlah:  constUangHarianJumlah,
			}
			personData[assignee.EmployeeID] = data
		}

		for _, tx := range assignee.Transactions {
			switch transaction.TransactionType(strings.ToLower(tx.Type)) {
			case transaction.TransactionTypeAccommodation:
				if tx.PaymentType == "uang_muka" {
					if tx.TotalNight != nil {
						data.UMPenginapanJmlHari += *tx.TotalNight
					}
					data.UMPenginapanPerhari = tx.Amount // Assuming Amount is per night
					data.UMPenginapanJumlah += tx.Subtotal
				} else {
					if tx.TotalNight != nil {
						data.RPenginapanJmlHari += *tx.TotalNight
					}
					data.RPenginapanPerhari = tx.Amount // Assuming Amount is per night
					data.RPenginapanJumlah += tx.Subtotal
				}
			case transaction.TransactionTypeTransport:
				if tx.PaymentType == "uang_muka" {
					if strings.ToLower(tx.Subtype) == "flight" {
						data.UMTransportTiketPesawat += tx.Subtotal
					}
					data.UMTransportJumlah += tx.Subtotal
				} else {
					if strings.ToLower(tx.Subtype) == "flight" {
						if data.UMTransportTiketPesawat == 0 {
							data.RTransportTiketPesawat += tx.Subtotal
						} else {
							data.RTransportTiketPesawat = data.UMTransportTiketPesawat
						}
					}

					if strings.ToLower(tx.Subtype) == "taxi" {
						if strings.ToLower(tx.TransportDetail) == "transport_darat" {
							data.RTransportDarat += tx.Subtotal
						}

						if strings.ToLower(tx.TransportDetail) == "transport_asal" {
							data.RTransportAsal += tx.Subtotal
						}

						if strings.ToLower(tx.TransportDetail) == "transport_daerah" {
							data.RTransportDaerah += tx.Subtotal
						}
					}
					data.RTransportJumlah += tx.Subtotal
				}
			case transaction.TransactionTypeOther:
				if tx.PaymentType == "uang_muka" {
					data.UMTotalDibayarkan += tx.Subtotal
				} else {
					data.RTotalDibayarkan += tx.Subtotal
				}
			}

			data.UMTotalDibayarkan = data.UMUangHarianJumlah + data.UMPenginapanJumlah + data.UMTransportJumlah
		}
	}

	personNo := 1
	for _, data := range personData {
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("A%d", currentRow), personNo); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("B%d", currentRow), data.Name); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("C%d", currentRow), data.NIP); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("D%d", currentRow), data.Jabatan); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("E%d", currentRow), data.Gol); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("F%d", currentRow), data.Tujuan); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("G%d", currentRow), data.Tanggal); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("H%d", currentRow), data.UMUangHarianJmlHari); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("I%d", currentRow), "Hari"); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("J%d", currentRow), data.UMUangHarianPerhari); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("K%d", currentRow), data.UMUangHarianJumlah); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("L%d", currentRow), data.UMPenginapanJmlHari); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("M%d", currentRow), "Hari"); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("N%d", currentRow), data.UMPenginapanPerhari); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("O%d", currentRow), data.UMPenginapanJumlah); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("P%d", currentRow), data.UMTransportTiketPesawat); err != nil {
			return currentRow, err
		}

		if sheetName == "PEMANTAUAN REKAP RAMPUNG" {
			if err := g.file.SetCellValue(sheetName, fmt.Sprintf("Q%d", currentRow), data.RTransportTiketPesawat); err != nil {
				return currentRow, err
			}
			if err := g.file.SetCellValue(sheetName, fmt.Sprintf("R%d", currentRow), data.RTransportAsal); err != nil {
				return currentRow, err
			}
			if err := g.file.SetCellValue(sheetName, fmt.Sprintf("S%d", currentRow), data.RTransportDaerah); err != nil {
				return currentRow, err
			}
			if err := g.file.SetCellValue(sheetName, fmt.Sprintf("T%d", currentRow), data.RTransportDarat); err != nil {
				return currentRow, err
			}
		} else {
			if err := g.file.SetCellValue(sheetName, fmt.Sprintf("Q%d", currentRow), "-"); err != nil {
				return currentRow, err
			}
			if err := g.file.SetCellValue(sheetName, fmt.Sprintf("R%d", currentRow), "-"); err != nil {
				return currentRow, err
			}
			if err := g.file.SetCellValue(sheetName, fmt.Sprintf("S%d", currentRow), "-"); err != nil {
				return currentRow, err
			}
		}
		if err := g.file.SetCellFormula(sheetName, fmt.Sprintf("T%d", currentRow), fmt.Sprintf("=P%d+Q%d+R%d+S%d", currentRow, currentRow, currentRow, currentRow)); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("U%d", currentRow), data.UMTotalDibayarkan); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellValue(sheetName, fmt.Sprintf("V%d", currentRow), data.NoSpd); err != nil {
			return currentRow, err
		}
		if err := g.file.SetCellStyle(sheetName, fmt.Sprintf("A%d", currentRow), fmt.Sprintf("V%d", currentRow), dataStyle); err != nil {
			return currentRow, err
		}

		currentRow++
		personNo++
	}
	return currentRow, nil
}

func (g *Generator) generateSummaryRow(sheetName string, currentRow int) error {
	summaryStyle, err := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "right",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Bold:   true,
			Size:   10,
			Family: "Tahoma",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		return err
	}

	totalRow := currentRow
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("B%d", totalRow), "JUMLAH"); err != nil {
		return err
	}

	if err := g.file.SetCellFormula(sheetName, fmt.Sprintf("J%d", totalRow), fmt.Sprintf("=SUM(J10:J%d)", currentRow-1)); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, fmt.Sprintf("N%d", totalRow), fmt.Sprintf("=SUM(N10:N%d)", currentRow-1)); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, fmt.Sprintf("T%d", totalRow), fmt.Sprintf("=SUM(T10:T%d)", currentRow-1)); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, fmt.Sprintf("U%d", totalRow), fmt.Sprintf("=SUM(U10:Y%d)", currentRow-1)); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, fmt.Sprintf("A%d", totalRow), fmt.Sprintf("V%d", totalRow), summaryStyle); err != nil {
		return err
	}
	return nil
}

func (g *Generator) generateSignatureRow(sheetName string, currentRow int) error {
	nameStyle, _ := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical: "center",
		},
		Font: &excelize.Font{
			Bold:   true,
			Size:   10,
			Family: "Tahoma",
		},
	})

	dataStyle, _ := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical: "center",
		},
		Font: &excelize.Font{
			Size:   10,
			Family: "Tahoma",
		},
	})

	totalRow := currentRow + 3
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("B%d", totalRow), "Mengetahui/Menyetujui"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("D%d", totalRow), "Setuju/Lunas dibayar"); err != nil {
		return err
	}

	totalRow += 1
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("B%d", totalRow), "Pejabat Pembuat Komitmen II"); err != nil {
		return err
	}

	currentDate := time.Now().Format("2 January 2006")
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("D%d", totalRow), fmt.Sprintf("Tanggal: %s", currentDate)); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("I%d", totalRow), "Yang Membayarkan"); err != nil {
		return err
	}

	totalRow += 1
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("B%d", totalRow), fmt.Sprintf("Unit Kerja %s", workUnit)); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("D%d", totalRow), "Bendahara Pengeluaran"); err != nil {
		return err
	}

	totalRow += 5
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("B%d", totalRow), commitmentOfficer); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("D%d", totalRow), expenditureTreasurer); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("I%d", totalRow), payer); err != nil {
		return err
	}

	totalRow += 1
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("B%d", totalRow), commitmentOfficerId); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("D%d", totalRow), expenditureTreasurerId); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, fmt.Sprintf("I%d", totalRow), payerId); err != nil {
		return err
	}

	err := g.file.SetCellStyle(sheetName, fmt.Sprintf("B%d", currentRow+2), fmt.Sprintf("L%d", totalRow), dataStyle)
	if err != nil {
		return err
	}

	err = g.file.SetCellStyle(sheetName, fmt.Sprintf("B%d", totalRow-1), fmt.Sprintf("L%d", totalRow-1), nameStyle)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) setColumnWidths(sheetName string) error {
	if err := g.file.SetColWidth(sheetName, "A", "A", 5); err != nil {
		return err
	} // No
	if err := g.file.SetColWidth(sheetName, "B", "B", 20); err != nil {
		return err
	} // Nama
	if err := g.file.SetColWidth(sheetName, "C", "C", 20); err != nil {
		return err
	} // Employee ID
	if err := g.file.SetColWidth(sheetName, "D", "D", 25); err != nil {
		return err
	} // Position
	if err := g.file.SetColWidth(sheetName, "E", "E", 25); err != nil {
		return err
	} // Rank
	if err := g.file.SetColWidth(sheetName, "F", "F", 30); err != nil {
		return err
	} // Tujuan
	if err := g.file.SetColWidth(sheetName, "G", "G", 25); err != nil {
		return err
	} // Tanggal

	if err := g.file.SetColWidth(sheetName, "H", "H", 10); err != nil {
		return err
	} // Uang Harian Jml Hari
	if err := g.file.SetColWidth(sheetName, "I", "I", 15); err != nil {
		return err
	} // Uang Harian Perhari
	if err := g.file.SetColWidth(sheetName, "J", "J", 15); err != nil {
		return err
	} // Uang Harian Jumlah

	if err := g.file.SetColWidth(sheetName, "L", "L", 10); err != nil {
		return err
	} // Penginapan Jml Hari
	if err := g.file.SetColWidth(sheetName, "M", "M", 15); err != nil {
		return err
	} // Penginapan Perhari
	if err := g.file.SetColWidth(sheetName, "N", "N", 15); err != nil {
		return err
	} // Penginapan Jumlah

	if err := g.file.SetColWidth(sheetName, "P", "P", 15); err != nil {
		return err
	} // Transport Tiket Pesawat
	if err := g.file.SetColWidth(sheetName, "Q", "S", 15); err != nil {
		return err
	} // Transport Asal, Daerah, Darat (Q,R,S)
	if err := g.file.SetColWidth(sheetName, "T", "T", 15); err != nil {
		return err
	} // Transport Jumlah

	if err := g.file.SetColWidth(sheetName, "U", "U", 20); err != nil {
		return err
	} // Jumlah Dibayarkan (Rp)
	return nil
}

func (g *Generator) generateKw(sheetName string, req dto.RecapReportDTO) error {
	if _, err := g.file.NewSheet(sheetName); err != nil {
		return err
	}

	titleStyle, _ := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Bold:   true,
			Size:   10,
			Family: "Tahoma",
		},
	})

	if err := g.file.SetCellValue(sheetName, "A1", "KEMENTERIAN KESEHATAN REPUBLIK INDONESIA"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A2", "DIREKTORAT JENDERAL"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A3", "PENANGGULANAN PENYAKIT"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A5", "J A K A R T A"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "A1", "L1"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "A2", "L2"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "A3", "L3"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "A5", "L5"); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "A1", "L5", titleStyle); err != nil {
		return err
	}

	basicStyle, _ := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical: "center",
		},
		Font: &excelize.Font{
			Size:   10,
			Family: "Tahoma",
		},
	})

	if err := g.file.SetCellValue(sheetName, "M1", "Tahun Anggaran"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "M2", "No Bukti"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "M3", "Akun"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "O1", ":"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O2", ":"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O3", ":"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "P1", "2025"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "P3", "024.05.WA.4815.EBD.953."); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "P4", "501.B.524111"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "T1", 1); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "M1", "P4", basicStyle); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "M1", "P4", basicStyle); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A6", "RINCIAN BIAYA PERJALANAN DINAS"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "A6", "S6"); err != nil {
	}
	if err := g.file.SetCellStyle(sheetName, "A6", "S6", titleStyle); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A7", "Lampiran SPD Nomor"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "F7", ":"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "G7", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AJ$100,22,FALSE)"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A8", "Tanggal"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "F8", ":"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "G8", req.StartDate); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "A7", "S9", basicStyle); err != nil {
		return err
	}

	if err := g.file.SetColWidth(sheetName, "A", "A", 4); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "B", "B", 3.4); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "C", "C", 4.5); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "D", "D", 5.17); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "E", "E", 1.50); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "F", "F", 6.00); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "G", "G", 1.17); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "H", "H", 12.33); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "I", "I", 2.83); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "J", "J", 2.67); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "K", "K", 2.83); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "L", "L", 4.00); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "M", "M", 11.00); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "N", "N", 2.83); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "O", "O", 2.67); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "P", "P", 19.17); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "Q", "Q", 3.50); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "R", "R", 1.17); err != nil {
		return err
	}
	if err := g.file.SetColWidth(sheetName, "S", "S", 3.67); err != nil {
		return err
	}

	// header
	if err := g.file.SetCellValue(sheetName, "A10", "NO"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "C10", "PERINCIAN BIAYA"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "L10", "JUMLAH"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O10", "KETERANGAN"); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "A10", "B11"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "C10", "K11"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "L10", "N11"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "O10", "S11"); err != nil {
		return err
	}

	kwHeaderStyle, _ := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Bold:   true,
			Size:   10,
			Family: "Tahoma",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 2},
			{Type: "top", Color: "000000", Style: 2},
			{Type: "bottom", Color: "000000", Style: 2},
			{Type: "right", Color: "000000", Style: 2},
		},
	})

	if err := g.file.SetCellStyle(sheetName, "A10", "B11", kwHeaderStyle); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "C10", "K11", kwHeaderStyle); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "L10", "N11", kwHeaderStyle); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "O10", "S11", kwHeaderStyle); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A13", "1"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "C13", "Uang harian :"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "C14", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$100,8,FALSE)"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "D14", "hr"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "E14", "x"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "F14", "Rp."); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "H14", "-"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "L14", "Rp."); err != nil {
		return err
	}

	if err := g.file.SetCellFormula(sheetName, "M14", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$12,11,FALSE)"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A15", "2"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "C15", "Transport"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "C16", "a."); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "D16", "Tiket :"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "D17", fmt.Sprintf("- Pesawat Jakarta - %s (PP)", req.DestinationCity)); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "L17", "Rp."); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "M17", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$100,20,FALSE)"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "C18", "b."); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "D18", "Transport (PP):"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "D19", "Transport Jakarta - Bandara Soetta (PP)"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "L19", "Rp."); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "M19", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$100,17,FALSE)"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "D20", fmt.Sprintf("Transport Dareah %s (PP)", req.DestinationCity)); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "L20", "Rp."); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "M20", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$100,18,FALSE)"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A21", "3"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "C21", "Biaya Penginapan : "); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "C22", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$100,12,FALSE)"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "D22", "==VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$100,13,FALSE)"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "D22", "mlm"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "E22", "x"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "F22", "Rp."); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "H22", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$100,14,FALSE)"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "L22", "Rp."); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "M22", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$100,15,FALSE)"); err != nil {
		return err
	}

	dynamicStyle := func(borderTypes []string, bold bool, italic bool, borderStyle int, horizontal string) int {
		var borders []excelize.Border

		for _, side := range borderTypes {
			switch side {
			case "top":
				borders = append(borders, excelize.Border{
					Type:  "top",
					Color: "000000",
					Style: borderStyle,
				})
			case "bottom":
				borders = append(borders, excelize.Border{
					Type:  "bottom",
					Color: "000000",
					Style: borderStyle,
				})
			case "left":
				borders = append(borders, excelize.Border{
					Type:  "left",
					Color: "000000",
					Style: borderStyle,
				})
			case "right":
				borders = append(borders, excelize.Border{
					Type:  "right",
					Color: "000000",
					Style: borderStyle,
				})
			}
		}

		style, _ := g.file.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: horizontal,
				Vertical:   "center",
			},
			Font: &excelize.Font{
				Italic: italic,
				Bold:   bold,
				Size:   10,
				Family: "Tahoma",
			},
			Border: borders,
		})

		return style
	}

	if err := g.file.SetCellStyle(sheetName, "C12", "C24", dynamicStyle([]string{"left"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "L12", "L24", dynamicStyle([]string{"left"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "O12", "O24", dynamicStyle([]string{"left"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "S12", "S24", dynamicStyle([]string{"right"}, false, false, 2, "left")); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "A13", "B13"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "A15", "B15"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "A21", "B21"); err != nil {
		return err
	}

	noStyle, _ := g.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "left",
		},
		Font: &excelize.Font{
			Size:   10,
			Family: "Tahoma",
		},
	})

	if err := g.file.SetCellStyle(sheetName, "A13", "A21", noStyle); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "C24", "J U M L A H"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "C24", "K24"); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "C24", "C24", kwHeaderStyle); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "A24", "B24", dynamicStyle([]string{"top", "bottom"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "D24", "K24", dynamicStyle([]string{"top", "bottom"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "M24", "N24", dynamicStyle([]string{"top", "bottom"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "L24", "L24", dynamicStyle([]string{"top", "bottom", "left"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "P24", "R24", dynamicStyle([]string{"top", "bottom"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "O24", "O24", dynamicStyle([]string{"top", "bottom", "left"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "S24", "S24", dynamicStyle([]string{"top", "bottom", "right"}, false, false, 2, "left")); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "L24", "Rp."); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "M24", "=SUM(M12:M23)"); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "M24", "M24", dynamicStyle([]string{"top", "bottom"}, true, false, 2, "left")); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "L24", "L24", dynamicStyle([]string{"top", "bottom", "left"}, true, false, 2, "left")); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A25", "TERBILANG"); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "A25", "A25", dynamicStyle([]string{}, false, true, 2, "left")); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "O27", fmt.Sprintf("Jakarta, %s", req.ReceiptSignatureDate)); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A28", "Telah dibayar sejumlah"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O28", "Telah menerima jumlah uang sebesar"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A29", "Rp."); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "C29", "=M24"); err != nil {
		return err
	}

	if err := g.file.SetCellFormula(sheetName, "O29", "=M24"); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "C29", "F29"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "O29", "P29"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A30", "Bendahara Pengeluaran Pembantu"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "I30", "PUM Timker"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O30", "Yang Menerima"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A31", "Unit Kerja Setditjen Penanggulangan Penyakit"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A34", expenditureTreasurer); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "I34", payer); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A35", "NIP"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "B35", expenditureTreasurerId); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "I34", payer); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "I35", "NIP"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "J35", payerId); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "O34", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$100,2,FALSE)"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O35", "NIP"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "P35", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$100,3,FALSE)"); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "A28", "S35", dynamicStyle([]string{}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "A28", "S35", dynamicStyle([]string{}, false, false, 2, "left")); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "A29", "C29", dynamicStyle([]string{}, true, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "O29", "O29", dynamicStyle([]string{}, true, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "A35", "S35", dynamicStyle([]string{"bottom"}, false, false, 2, "left")); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A36", "PERHITUNGAN SPD UANG MUKA"); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "A36", "A36", dynamicStyle([]string{}, true, false, 0, "center")); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "A36", "S36"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A37", "Ditetapkan sejumlah"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "I37", ":"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "J37", "Rp"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "K37", "=M24"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "K37", "M37"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A38", "Yang telah dibayar semula"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "I38", ":"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "J38", "Rp"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "K38", "=K37"); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "K38", "M38"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A39", "Sisa Kurang/Lebih"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "I39", ":"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "J39", "Rp"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "K39", "=K37"); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "K39", "M39"); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "A37", "L40", dynamicStyle([]string{}, true, false, 0, "left")); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "M40", "an.  Kuasa Pengguna Anggaran"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "M41", "      Pejabat Pembuat Komitmen,"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "M42", "      Satker Kantor Pusat Ditjen Penanggulangan Penyakit"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "M43", "      Unit Kerja Setditjen Penanggulangan Penyakit"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "M47", fmt.Sprintf("      %s", expenditureTreasurer)); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "M48", fmt.Sprintf("      NIP  %s", expenditureTreasurer)); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "A50", "S50", dynamicStyle([]string{"bottom"}, false, false, 8, "left")); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A53", "DAFTAR PENELUARAN RIIL"); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "A53", "S53"); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "A53", "A53", dynamicStyle([]string{}, true, false, 0, "center")); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A55", "Yang bertanda tangan dibawah ini  :"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A57", "Nama"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "C57", ":"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "D57", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$12,2,FALSE)"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A58", "NIP"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "C58", ":"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "D58", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$12,3,FALSE)"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A59", "Jabatan"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "C59", ":"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "D59", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$12,4,FALSE)"); err != nil {
		return err
	}

	if err := g.file.SetCellFormula(sheetName, "A61", `="Berdasarkan Surat Perjalanan Dinas ( SPD ) Nomor "&VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AJ$100,22,FALSE)`); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "N61", fmt.Sprintf("tanggal %s", req.SpdDate)); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A62", "dengan ini kami menyatakan dengan sesungguhnya bahwa :"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A64", "1.  Biaya transport pegawai dan / atau biaya penginapan dibawah ini yang tidak dapat "); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A65", "     diperoleh bukti-bukti pengeluarannya, meliputi  :"); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A67", "No"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "D67", "U r a i a n"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O67", "Jumlah"); err != nil {
		return err
	}

	if err := g.file.MergeCell(sheetName, "A67", "C67"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "D67", "N67"); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "O67", "S67"); err != nil {
		return err
	}

	if err := g.file.SetCellStyle(sheetName, "A67", "S67", dynamicStyle([]string{"top", "right", "bottom", "left"}, true, false, 2, "center")); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "B69", "1"); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "B69", "B69", dynamicStyle([]string{}, false, false, 0, "center")); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "D69", "Transport :"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "E70", "=D19"); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "E71", "=D20"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O70", "Rp."); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O71", "Rp."); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "D74", "JUMLAH"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O74", "Rp."); err != nil {
		return err
	}
	if err := g.file.MergeCell(sheetName, "D74", "N74"); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "D68", "D73", dynamicStyle([]string{"left"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "O68", "O73", dynamicStyle([]string{"left"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "S68", "S73", dynamicStyle([]string{"right"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "A74", "S74", dynamicStyle([]string{"top", "bottom"}, false, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "D74", "D74", dynamicStyle([]string{"top", "bottom", "left"}, true, false, 2, "center")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "O74", "O74", dynamicStyle([]string{"top", "bottom", "left"}, true, false, 2, "left")); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "S74", "S74", dynamicStyle([]string{"top", "bottom", "right"}, false, false, 2, "left")); err != nil {
		return err
	}

	if err := g.file.SetCellValue(sheetName, "A76", "2.  Jumlah uang tersebut pada angka 1 diatas benar-benar dikeluarkan untuk pelaksanaan "); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A77", "     perjalanan dinas dimaksud dan apabila kemudian hari terdapat kelebihan atas pembayaran"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A78", "     kami bersedia untuk menyetorkan kelebihan tersebut ke Kas Negara."); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A80", "Demikian pernyataan ini kami buat dengan sebenarnya, untuk dipergunakan sebagaimana mestinya."); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A83", "Mengetahui / Menyetujui"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "N83", fmt.Sprintf("Jakarta, %s", req.ReceiptSignatureDate)); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "N84", "Pelaksana SPD,"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A85", "Satker Kantor Pusat Ditjen Penanggulangan Penyakit"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A86", "Unit Kerja Setditjen Penanggulangan Penyakit"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A89", expenditureTreasurer); err != nil {
		return err
	}
	if err := g.file.SetCellFormula(sheetName, "N89", "=VLOOKUP($T$1,'PEMANTAUAN REKAP UANG MUKA'!$A$11:$AC$100,2,FALSE)"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "A90", "NIP"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "B90", expenditureTreasurerId); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "N90", "NIP"); err != nil {
		return err
	}
	if err := g.file.SetCellValue(sheetName, "O90", expenditureTreasurerId); err != nil {
		return err
	}
	if err := g.file.SetCellStyle(sheetName, "A89", "S89", dynamicStyle([]string{}, true, false, 0, "left")); err != nil {
		return err
	}

	// dengan ini kami menyatakan dengan sesungguhnya bahwa :

	return nil
}

// GenerateRecapExcel creates an Excel file based on transaction data
func (g *Generator) GenerateRecapExcel(req dto.RecapReportDTO) (*bytes.Buffer, error) {
	var err error
	sheetName := "PEMANTAUAN REKAP UANG MUKA"
	// Remove the default sheet created by NewFile
	g.file.DeleteSheet("Sheet1")
	if err = g.file.SetSheetName(g.file.GetSheetName(0), sheetName); err != nil {
		return nil, err
	}

	if err = g.generateTitle(sheetName); err != nil {
		return nil, err
	}

	if err = g.generateTableHeader(sheetName); err != nil {
		return nil, err
	}

	// Data starting row
	currentRow := 11
	if currentRow, err = g.generateTableData(sheetName, req.Assignees, currentRow); err != nil {
		return nil, err
	}

	if err = g.generateSummaryRow(sheetName, currentRow); err != nil {
		return nil, err
	}

	if err = g.generateSignatureRow(sheetName, currentRow); err != nil {
		return nil, err
	}

	if err = g.setColumnWidths(sheetName); err != nil {
		return nil, err
	}

	kwUangMuka := "KW UANG MUKA"

	err = g.generateKw(kwUangMuka, req)
	if err != nil {
		return nil, err
	}

	sheetName = "PEMANTAUAN REKAP RAMPUNG"
	if _, err = g.file.NewSheet(sheetName); err != nil {
		return nil, err
	}

	if err = g.generateTitle(sheetName); err != nil {
		return nil, err
	}

	if err = g.generateTableHeader(sheetName); err != nil {
		return nil, err
	}

	if currentRow, err = g.generateTableData(sheetName, req.Assignees, 11); err != nil {
		return nil, err
	}

	if err = g.generateSummaryRow(sheetName, currentRow); err != nil {
		return nil, err
	}

	if err = g.generateSignatureRow(sheetName, currentRow); err != nil {
		return nil, err
	}

	if err = g.setColumnWidths(sheetName); err != nil {
		return nil, err
	}

	if _, err = g.file.NewSheet("KW RAMPUNG"); err != nil {
		return nil, err
	}

	// Save the Excel file to a buffer
	var b bytes.Buffer
	if err = g.file.Write(&b); err != nil {
		return nil, fmt.Errorf("failed to write excel to buffer: %w", err)
	}

	return &b, nil
}

var angka = []string{
	"", "Satu", "Dua", "Tiga", "Empat", "Lima", "Enam", "Tujuh", "Delapan", "Sembilan", "Sepuluh", "Sebelas",
}

func Terbilang(n int) string {
	if n < 12 {
		return angka[n]
	} else if n < 20 {
		return Terbilang(n-10) + " Belas"
	} else if n < 100 {
		return Terbilang(n/10) + " Puluh " + Terbilang(n%10)
	} else if n < 200 {
		return "Seratus " + Terbilang(n-100)
	} else if n < 1000 {
		return Terbilang(n/100) + " Ratus " + Terbilang(n%100)
	} else if n < 2000 {
		return "Seribu " + Terbilang(n-1000)
	} else if n < 1000000 {
		return Terbilang(n/1000) + " Ribu " + Terbilang(n%1000)
	} else if n < 1000000000 {
		return Terbilang(n/1000000) + " Juta " + Terbilang(n%1000000)
	} else if n < 1000000000000 {
		return Terbilang(n/1000000000) + " Miliar " + Terbilang(n%1000000000)
	} else if n < 1000000000000000 {
		return Terbilang(n/1000000000000) + " Triliun " + Terbilang(n%1000000000000)
	}
	return ""
}

func CleanSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
