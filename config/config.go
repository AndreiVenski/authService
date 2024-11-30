package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Postgres PostgresqlConfig
	Server   ServerConfig
}

type PostgresqlConfig struct {
	PostgresqlHost     string `envconfig:"POSTGRESQL_HOST"`
	PostgresqlPort     string `envconfig:"POSTGRESQL_PORT"`
	PostgresqlUser     string `envconfig:"POSTGRESQL_USER"`
	PostgresqlDbname   string `envconfig:"POSTGRESQL_DBNAME"`
	PostgresqlPassword string `envconfig:"POSTGRESQL_PASSWORD"`
}

type ServerConfig struct {
	RunningPort                string `envconfig:"SERVER_RUNNINGPORT"`
	JWTSecret                  string `envconfig:"SERVER_JWTSECRET"`
	AccessTokenExpiresHourInt  int    `envconfig:"SERVER_ACCESSTOKENEXPIRESHOURINT"`
	RefreshTokenExpiresHourInt int    `envconfig:"SERVER_REFRESHTOKENEXPIRESHOURINT"`
}

func InitConfig(path string) (*Config, error) {
	var cfg Config
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	if err = envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
