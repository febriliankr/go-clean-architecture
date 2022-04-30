package entity

import "time"

type Truck struct {
	ID             int64      `json:"id" db:"id"`
	LicenseNumber  string     `json:"license_number" db:"license_number"`
	TruckType      string     `json:"truck_type" db:"truck_type"`
	PlateType      string     `json:"plate_type" db:"plate_type"`
	ProductionYear int64      `json:"production_year" db:"production_year"`
	Status         bool       `json:"status" db:"status"`
	ShipmentStatus bool       `json:"shipment_status" db:"shipment_status"`
	STNK           string     `json:"stnk" db:"stnk"`
	KIR            string     `json:"kir" db:"kir"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at" db:"deleted_at"`
}

// Requests
type (
	CreateNewTruckRequest struct {
		LicenseNumber  string `json:"license_number" validate:"required"`
		TruckType      string `json:"truck_type" validate:"required"`
		PlateType      string `json:"plate_type" validate:"required"`
		ProductionYear int64  `json:"production_year" validate:"required"`
		STNK           string `json:"stnk" db:"stnk"`
		KIR            string `json:"kir" db:"kir"`
	}

	GetTruckListRequest struct {
		LicenseNumber string `json:"license"`
		TruckType     string `json:"type"`
		Page          int64  `json:"page"`
		Limit         int64  `json:"limit"`
	}

	GetTruckByIDRequest struct {
		ID int64 `json:"id"`
	}

	UpdateTruckRequest struct {
		ID             int64  `json:"id"  validate:"required"`
		LicenseNumber  string `json:"license_number"`
		TruckType      string `json:"truck_type"`
		PlateType      string `json:"plate_type"`
		ProductionYear int64  `json:"production_year"`
		Status         bool   `json:"status"`
		STNK           string `json:"stnk"`
		KIR            string `json:"kir"`
	}

	DeleteTruckRequest struct {
		ID int64 `json:"id"  validate:"required"`
	}
)

// Response
type (
	CreateNewTruckResponse struct {
		ID int64 `json:"id"`
	}

	GetTruckListResponse struct {
		Data struct {
			Rows []Truck `json:"rows"`
		} `json:"data"`
	}

	GetTruckByIDResponse struct {
		Truck
	}

	UpdateTruckResponse struct {
	}

	DeleteTruckResponse struct {
	}
)
