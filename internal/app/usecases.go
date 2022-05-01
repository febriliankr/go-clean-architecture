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
}

// Inject dependency for usecase layer
func newUsecases(cfg config.Config, repos *Repos) (*Usecases, error) {
	return &Usecases{
		DriverUC:   usecase.NewDriverUsecase(repos.DriverRepo),
		ShipmentUC: usecase.NewShipmentUsecase(repos.ShipmentRepo),
		TruckUC:    usecase.NewTruckUsecase(repos.TruckRepo),
	}, nil
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
