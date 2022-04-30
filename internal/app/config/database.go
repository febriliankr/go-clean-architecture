package config

type Database struct {
	Address      string `env:"DATABASE_URL"`
	Driver       string `env:"DATABASE_DRIVER"`
	MaxConns     int    `env:"DATABASE_MAX_CONNS"`
	MaxIdleConns int    `env:"DATABASE_MAX_IDLE"`
}
