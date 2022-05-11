package app

import (
	"github.com/febriliankr/go-clean-architecture/internal"
	"github.com/febriliankr/go-clean-architecture/internal/app/config"
	"github.com/febriliankr/go-clean-architecture/internal/usecase"
)

type Usecases struct {
	DriverUC   internal.DriverUC
	ShipmentUC internal.ShipmentUC
	TruckUC    internal.TruckUC
	PaymentUC  internal.PaymentUC
}

// Inject dependency for usecase layer
func newUsecases(cfg config.Config, repos *Repos) *Usecases {
	driverUC := usecase.NewDriverUsecase(repos.DriverRepo)
	uc := &Usecases{
		DriverUC:   driverUC,
		ShipmentUC: usecase.NewShipmentUsecase(repos.ShipmentRepo),
		TruckUC:    usecase.NewTruckUsecase(repos.TruckRepo),
		PaymentUC:  usecase.NewPaymentUsecase(repos.PaymentRepo),
	}
	return uc
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
