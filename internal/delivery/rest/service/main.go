package service

import (
	"github.com/febriliankr/go-clean-architecture/internal/app"
)

type Services struct {
	*DriverService
	*ShipmentService
	*TruckService
	*PaymentService
}

func GetServices(app *app.SeminarApp) *Services {
	return &Services{
		DriverService:   NewDriverService(app),
		ShipmentService: NewShipmentService(app),
		TruckService:    NewTruckService(app),
		PaymentService:  NewPaymentService(app),
	}
}
