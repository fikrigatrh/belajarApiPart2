package createPaymentRepository

type InputCreatePayment struct {
	IdUnit                  uint   `json:"id_unit" validate:"required,numeric"`
	DimintaOleh             string `json:"diminta_oleh" validate:"required"`
	Keperluan               string `json:"keperluan" validate:"required"`
	TanggalPembayaranAktual string `json:"tanggal_pembayaran_aktual" validate:"required"`
	JumlahPayment           int64  `json:"jumlah_payment" validate:"required,numeric"`
	Terbilang               string `json:"terbilang" validate:"required"`
	NamaRekPenerima         string `json:"nama_rek_penerima" validate:"required"`
	NoRekPenerima           string `json:"no_rek_penerima" validate:"required"`
	RequestTerkirim 		string `json:"request_terkirim"`
	Status 					int `json:"status"`
}
