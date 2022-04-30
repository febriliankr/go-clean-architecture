package config

type Server struct {
	Port                string `env:"PORT"`
	Debug               bool   `env:"SERVER_DEBUG"`
	ReadTimeoutSeconds  int    `env:"SERVER_READ_TIMEOUT"`
	WriteTimeoutSeconds int    `env:"SERVER_WRITE_TIMEOUT"`
}
