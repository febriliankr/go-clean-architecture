package app

import (
	"github.com/caarlos0/env/v6"
	"github.com/febriliankr/go-clean-architecture/internal/app/config"
	"github.com/pkg/errors"
)

type SeminarApp struct {
	Repos    *Repos
	Usecases *Usecases

	Cfg config.Config
}

// NewSeminarApp Initialize app and do dependency injection for all layers
func NewSeminarApp() (*SeminarApp, error) {
	cfg, err := readConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "error getting config")
	}

	app := new(SeminarApp)

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

func (a *SeminarApp) Close() []error {
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
