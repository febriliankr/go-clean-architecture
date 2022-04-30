package app

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"github.com/ridwanakf/tms-backend/internal/app/config"
)

type ShipmentApp struct {
	Repos    *Repos
	Usecases *Usecases

	Cfg config.Config
}

// NewShipmentApp Initialize app and do dependency injection for all layers
func NewShipmentApp() (*ShipmentApp, error) {
	cfg, err := readConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "error getting config")
	}

	app := new(ShipmentApp)

	app.Cfg = cfg

	app.Repos, err = newRepos(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "errors invoking newRepos")
	}

	app.Usecases, err = newUsecases(cfg, app.Repos)
	if err != nil {
		return nil, errors.Wrapf(err, "errors invoking newUsecases")
	}

	return app, nil
}

func (a *ShipmentApp) Close() []error {
	var errs []error

	errs = append(errs, a.Repos.Close()...)
	errs = append(errs, a.Usecases.Close()...)

	return errs
}

func readConfig() (config.Config, error) {
	var cfg config.Config

	// Read vars that exist in ENV
	if err := env.Parse(&cfg); err != nil {
		return config.Config{}, errors.Wrapf(err, "error reading config from ENV")
	}

	return cfg, nil
}
