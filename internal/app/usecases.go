package app

import (
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/app/config"
	"github.com/ridwanakf/tms-backend/internal/usecase"
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
