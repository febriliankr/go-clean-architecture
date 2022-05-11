package entity

type User struct {
	ID              int64  `json:"id" db:"id"`
	NamaLengkap     string `json:"nama_lengkap" db:"nama_lengkap"`
	Email           string `json:"email" db:"email"`
	NomorHP         string `json:"nomor_hp" db:"nomor_hp"`
	FirebaseID      string `json:"firebaseId" db:"firebaseId"`
	Institusi       string `json:"institusi" db:"institusi"`
	Pekerjaan       string `json:"pekerjaan" db:"pekerjaan"`
	CheckEmailPromo string `json:"check_email_promo" db:"check_email_promo"`
	DisplayName     string `json:"displayName" db:"displayName"`
	PhotoURL        string `json:"photoUrl" db:"photoUrl"`
	NoSTR           string `json:"no_str" db:"no_str"`
}

type (
	CreateNewUserRequest struct {
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}

	GetUserListRequest struct {
		Page  int64 `json:"page"`
		Limit int64 `json:"limit"`
	}

	GetUserByUUIDRequest struct {
		UUID string `json:"uuid"`
	}
)

type (
	GetUserListResponse struct {
		Data struct {
			Rows []User `json:"rows"`
		} `json:"data"`
	}
)

// id
// nama_lengkap
// email
// nomor_hp
// firebaseId
// institusi
// pekerjaan
// check_email_promo
// displayName
// photoUrl
// no_str
