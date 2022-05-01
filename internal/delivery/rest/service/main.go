package service

import (
	"github.com/febriliankr/go-clean-architecture/internal/app"
)

type Services struct {
	*DriverService
	*ShipmentService
	*TruckService
}

func GetServices(app *app.ShipmentApp) *Services {
	return &Services{
		DriverService:   NewDriverService(app),
		ShipmentService: NewShipmentService(app),
		TruckService:    NewTruckService(app),
	}
}
