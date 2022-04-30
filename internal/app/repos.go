package app

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/app/config"
	"github.com/ridwanakf/tms-backend/internal/repo"
)

type Repos struct {
	DriverRepo   internal.DriverRepo
	ShipmentRepo internal.ShipmentRepo
	TruckRepo    internal.TruckRepo
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
	}, nil
}

func initDB(cfg config.Config) (*sqlx.DB, error) {
	// Connect SQL DB

	log.Println(cfg.DB.Driver)
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
