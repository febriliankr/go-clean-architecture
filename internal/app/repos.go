package app

import (
	"github.com/febriliankr/go-clean-architecture/internal"
	"github.com/febriliankr/go-clean-architecture/internal/app/config"
	"github.com/febriliankr/go-clean-architecture/internal/repo"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repos struct {
	DriverRepo   internal.DriverRepo
	ShipmentRepo internal.ShipmentRepo
	TruckRepo    internal.TruckRepo
	PaymentRepo  internal.PaymentRepo
}

// Inject dependency for repository layer
func newRepos(cfg config.Config) (*Repos, error) {
	db, err := initDB(cfg)
	if err != nil {
		return nil, err
	}
	return &Repos{
		DriverRepo:   repo.NewDriverDB(db),
		ShipmentRepo: repo.NewShipmentDB(db),
		TruckRepo:    repo.NewTruckDB(db),
		PaymentRepo:  repo.NewPaymentDB(db),
	}, nil
}

func initDB(cfg config.Config) (*sqlx.DB, error) {
	// Connect SQL DB
	db, err := sqlx.Connect(cfg.DB.Driver, cfg.DB.Address)
	if err != nil {
		return nil, err
	}

	// Set db params
	db.SetMaxIdleConns(cfg.DB.MaxConns)
	db.SetMaxOpenConns(cfg.DB.MaxIdleConns)

	return db, nil
}

func (r *Repos) Close() []error {
	var errs []error

	return errs
}
