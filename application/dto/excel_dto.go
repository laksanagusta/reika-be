package dto

type TravelExpenseReportRowDTO struct {
	No                int    `json:"no"`
	Name              string `json:"name"`
	NIP               string `json:"nip"`
	Jabatan           string `json:"jabatan"`
	Gol               string `json:"gol"`
	Tujuan            string `json:"tujuan"`
	Tanggal           string `json:"tanggal"`
	UangHarianJmlHari int32  `json:"uang_harian_jml_hari"`
	UangHarianPerhari int32  `json:"uang_harian_perhari"`
	UangHarianJumlah  int32  `json:"uang_harian_jumlah"`
	PenginapanJmlHari int32  `json:"penginapan_jml_hari"`
	PenginapanPerhari int32  `json:"penginapan_perhari"`
	PenginapanJumlah  int32  `json:"penginapan_jumlah"`
	TiketPesawat      int32  `json:"tiket_pesawat"`
	TransportAsal     int32  `json:"transport_asal"`
	TransportDaerah   int32  `json:"transport_daerah"`
	TransportDarat    int32  `json:"transport_darat"`
	TransportJumlah   int32  `json:"transport_jumlah"`
	JumlahDibayarkan  int32  `json:"jumlah_dibayarkan"`
}

// GenerateExcelReportRequest represents the request to generate an Excel report.
type GenerateExcelReportRequest struct {
	Data []TravelExpenseReportRowDTO `json:"data"`
}

// GenerateExcelReportResponse represents the response containing the generated Excel file.
type GenerateExcelReportResponse struct {
	FileContent []byte `json:"file_content"`
	FileName    string `json:"file_name"`
}
