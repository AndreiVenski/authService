package server

import (
	"authService/config"
	loggerimp "authService/pkg/logger"
	"github.com/jmoiron/sqlx"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	db     *sqlx.DB
	http   *fiber.App
	cfg    *config.Config
	logger loggerimp.Logger
}

func NewServer(cfg *config.Config, http *fiber.App, logger loggerimp.Logger, db *sqlx.DB) *Server {
	return &Server{
		db:     db,
		http:   http,
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) Run() error {
	return nil
}
