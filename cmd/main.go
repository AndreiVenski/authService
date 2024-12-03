package main

import (
	_ "authService/api/docs"
	"authService/config"
	"authService/internal/emailService/mocks"
	server2 "authService/internal/server"
	"authService/pkg/db/postgres_conn"
	"authService/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/pressly/goose/v3"
	"os"
)

// @title AuthService API
// @version 1.0
// @description This is API for AuthService
// @contact.name Andrei Venski
// @contact.url https://github.com/andrew967
// @contact.email venskiandrei32@gmail.com
// @BasePath /api/v1/auth
func main() {
	loggerApi := logger.NewApiLogger()
	loggerApi.InitLogger()

	cfg, err := config.InitConfig(".env")
	if err != nil {
		loggerApi.Error("Config init failed", err)
		os.Exit(1)
	}

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
