package config

import "github.com/kelseyhightower/envconfig"

type Server struct {
	Port int
}

type Postgres struct {
	Username string
	Password string
	Host     string
	Database string
	Port     int
}

type Broker struct {
	Username string
	Password string
	Host     string
	Port     int
}

type Config struct {
	DB   Postgres
	Srvr Server
	Brkr Broker
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Srvr); err != nil {
		return nil, err
	}

	if err := envconfig.Process("broker", &cfg.Brkr); err != nil {
		return nil, err
	}

	return cfg, nil
}
