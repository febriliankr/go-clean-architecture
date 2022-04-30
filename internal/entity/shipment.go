package entity

import "time"

type ShipmentStatus int64

const (
	Created ShipmentStatus = iota
	Allocated
	OngoingToOrigin
	AtOrigin
	OngoingToDestination
	AtDestination
	Completed
)

func (s ShipmentStatus) String() string {
	switch s {
	case Created:
		return "Created"
	case Allocated:
		return "Allocated"
	case OngoingToOrigin:
		return "Ongoing to Origin"
	case AtOrigin:
		return "At Origin"
	case OngoingToDestination:
		return "Ongoing to Destination"
	case AtDestination:
		return "At Destination"
	case Completed:
		return "Completed"
	}
	return "unknown"
}

type Shipment struct {
	ID             int64          `json:"id" db:"id"`
	ShipmentNumber string         `json:"no_shipment" db:"no_shipment"`
	DriverID       *int64         `json:"id_driver" db:"id_driver"`
	TruckID        *int64         `json:"id_truck" db:"id_truck"`
	Origin         string         `json:"origin" db:"origin"`
	Destination    string         `json:"destination" db:"destination"`
	Status         ShipmentStatus `json:"status" db:"status"`
	LoadingDate    time.Time      `json:"loading_date" db:"loading_date"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time     `json:"deleted_at" db:"deleted_at"`
}

// Requests
type (
	CreateNewShipmentRequest struct {
		ShipmentNumber string         `json:"no_shipment"`
		Origin         string         `json:"origin" validate:"required"`
		Destination    string         `json:"destination" validate:"required"`
		Status         ShipmentStatus `json:"status"`
		LoadingDateStr string         `json:"loading_date" validate:"required"`
		LoadingDate    time.Time
	}

	AllocateShipmentRequest struct {
		ID       int64 `json:"id" validate:"required"`
		DriverID int64 `json:"id_driver" validate:"required"`
		TruckID  int64 `json:"id_truck" validate:"required"`
	}

	GetShipmentListRequest struct {
	}

	GetShipmentByIDRequest struct {
	}

	UpdateShipmentRequest struct {
	}

	DeleteShipmentRequest struct {
	}
)

// Response
type (
	CreateNewShipmentResponse struct {
		ID int64 `json:"id"`
	}

	AllocateShipmentResponse struct {
	}

	GetShipmentListResponse struct {
	}

	GetShipmentByIDResponse struct {
	}

	UpdateShipmentResponse struct {
	}

	DeleteShipmentResponse struct {
	}
)
