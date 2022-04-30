package entity

type PaymentStatus int64

type Payment struct {
	ID                       int64  `json:"id" db:"id"`
	TotalPrice               string `json:"total_price" db:"total_price"`
	UniqueCode               int    `json:"unique_code" db:"unique_code"`
	StatusPembayaran         string `json:"status_pembayaran" db:"status_pembayaran"`
	JumlahTiket              int    `json:"jumlah_tiket" db:"jumlah_tiket"`
	UserEmail                string `json:"user_email" db:"user_email"`
	TanggalPembelian         string `json:"tanggal_pembelian" db:"tanggal_pembelian"`
	BuktiPembayaranUrl       string `json:"bukti_pembayaran_url" db:"bukti_pembayaran_url"`
	MetodePembayaran         string `json:"metode_pembayaran" db:"metode_pembayaran"`
	SeminarID                int    `json:"seminar_id" db:"seminar_id"`
	EventID                  int    `json:"event_id" db:"event_id"`
	PaymentReminderTimestamp string `json:"payment_reminder_timestamp" db:"payment_reminder_timestamp"`
	PaymentExpiredTimestamp  string `json:"payment_expired_timestamp" db:"payment_expired_timestamp"`
	EventReminderTimestamp   string `json:"event_reminder_timestamp" db:"event_reminder_timestamp"`
	BankAccountNumber        string `json:"bank_account_number" db:"bank_account_number"`
	BankAccountName          string `json:"bank_account_name" db:"bank_account_name"`
	Hidden                   bool   `json:"hidden" db:"hidden"`
}

// Requests
type (
	CreateNewPaymentRequest struct {
		UserEmail          string `json:"user_email" db:"user_email"`
		UniqueCode         string `json:"unique_code" db:"unique_code"`
		BuktiPembayaranUrl string `json:"bukti_pembayaran_url" db:"bukti_pembayaran_url"`
		MetodePembayaran   string `json:"metode_pembayaran" db:"metode_pembayaran"`
		BankAccountNumber  string `json:"bank_account_number" db:"bank_account_number"`
		BankAccountName    string `json:"bank_account_name" db:"bank_account_name"`
		Hidden             bool   `json:"hidden" db:"hidden"`
	}

	GetPaymentListRequest struct {
		Page  int64 `json:"page"`
		Limit int64 `json:"limit"`
	}

	GetPaymentByIDRequest struct {
		ID int64 `json:"id"`
	}

	UpdatePaymentRequest struct {
		ID                       int64  `json:"id"  validate:"required"`
		PaymentExpiredTimestamp  string `json:"payment_expired_timestamp" db:"payment_expired_timestamp"`
		PaymentReminderTimestamp string `json:"payment_reminder_timestamp" db:"payment_reminder_timestamp"`
		EventReminderTimestamp   string `json:"event_reminder_timestamp" db:"event_reminder_timestamp"`
		BuktiPembayaranUrl       string `json:"bukti_pembayaran_url" db:"bukti_pembayaran_url"`
		Hidden                   bool   `json:"hidden" db:"hidden"`
	}

	DeletePaymentRequest struct {
		ID int64 `json:"id"  validate:"required"`
	}
)

// Response
type (
	CreateNewPaymentResponse struct {
		ID int64 `json:"id"`
	}

	AllocatePaymentResponse struct {
	}

	GetPaymentListResponse struct {
		Data struct {
			Rows []Payment `json:"rows"`
		} `json:"data"`
	}

	GetPaymentByIDResponse struct {
	}

	UpdatePaymentResponse struct {
	}

	DeletePaymentResponse struct {
	}
)
