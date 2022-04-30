package entity

import "time"

type Driver struct {
	ID             int64      `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	Phone          string     `json:"phone" db:"phone"`
	CardID         int64      `json:"id_card" db:"id_card"`
	DriverLicense  int64      `json:"driver_license" db:"driver_license"`
	Status         bool       `json:"status" db:"status"`
	ShipmentStatus bool       `json:"shipment_status" db:"shipment_status"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at" db:"deleted_at"`
}

// Requests
type (
	CreateNewDriverRequest struct {
		Name           string `json:"name" db:"name"`
		Phone          string `json:"phone" db:"phone"`
		CardID         int64  `json:"id_card" db:"id_card"`
		DriverLicense  int64  `json:"driver_license" db:"driver_license"`
		Status         bool   `json:"status" db:"status"`
		ShipmentStatus bool   `json:"shipment_status" db:"shipment_status"`
	}

	GetDriverListRequest struct {
		Name           string `json:"name" db:"name"`
		Phone          string `json:"phone" db:"phone"`
		CardID         int64  `json:"id_card" db:"id_card"`
		DriverLicense  int64  `json:"driver_license" db:"driver_license"`
		Status         bool   `json:"status" db:"status"`
		ShipmentStatus bool   `json:"shipment_status" db:"shipment_status"`
		Page           int64  `json:"page"`
		Limit          int64  `json:"limit"`
	}

	GetDriverByIDRequest struct {
		ID int64 `json:"id"`
	}

	UpdateDriverRequest struct {
		ID             int64  `json:"id"`
		Name           string `json:"name" db:"name"`
		Phone          string `json:"phone" db:"phone"`
		CardID         int64  `json:"id_card" db:"id_card"`
		DriverLicense  int64  `json:"driver_license" db:"driver_license"`
		Status         bool   `json:"status" db:"status"`
		ShipmentStatus bool   `json:"shipment_status" db:"shipment_status"`
	}

	DeleteDriverRequest struct {
		ID int64 `json:"id"`
	}
)

// Response
type (
	CreateNewDriverResponse struct {
		ID int64 `json:"id"`
	}

	GetDriverListResponse struct {
		Data struct {
			Rows []Driver `json:"rows"`
		} `json:"data"`
	}

	GetDriverByIDResponse struct {
		Driver
	}

	UpdateDriverResponse struct {
	}

	DeleteDriverResponse struct {
	}
)
