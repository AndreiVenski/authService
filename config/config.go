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
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlDbname   string
	PostgresqlPassword string
}

type ServerConfig struct {
	RunningPort                string
	JWTSecret                  string
	AccessTokenExpiresHourInt  int
	RefreshTokenExpiresHourInt int
}

func InitConfig(path string) (*Config, error) {
	var cfg Config
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	err = envconfig.Process("", &cfg)
	return &cfg, err
}
