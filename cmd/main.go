package main

import (
	"authService/config"
	"authService/intern/emailService/mocks"
	server2 "authService/intern/server"
	"authService/pkg/db/postgres_conn"
	"authService/pkg/logger"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/pressly/goose/v3"
)

func main() {
	loggerApi := logger.NewApiLogger()
	loggerApi.InitLogger()

	cfg, err := config.InitConfig(".env")
	if err != nil {
		loggerApi.Error("Config init failed", err)
		os.Exit(1)
	}
	loggerApi.Info("DB_NAME:", cfg.Postgres.PostgresqlDbname, " username:", cfg.Postgres.PostgresqlUser)

	db, err := postgres_conn.NewPsqlDB(cfg)
	if err != nil {
		loggerApi.Error("DB init failed", err)
		os.Exit(1)
	}
	defer db.Close()

	if err = goose.Up(db.DB, "migrations"); err != nil {
		loggerApi.Error("Migrations up failed", err)
		os.Exit(1)
	}

	email := mocks.NewMockEmailService(loggerApi)

	fiberApp := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: false,
	})
	server := server2.NewServer(cfg, fiberApp, loggerApi, db, email)

	if err = server.Run(); err != nil {
		os.Exit(0)
	}
}
