package config

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
}
