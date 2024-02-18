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

type Mongo struct {
	Username string
	Password string
	Host     string
	Database string
	AuthDB   string
	Port     int
}

type Config struct {
	PostgresDB Postgres
	MongoDB    Mongo
	Srvr       Server
	Brkr       Broker
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("postgres", &cfg.PostgresDB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("mongo", &cfg.MongoDB); err != nil {
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
