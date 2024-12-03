package server

import (
	"authService/config"
	"authService/internal/emailService"
	loggerimp "authService/pkg/logger"
	"context"
	"github.com/jmoiron/sqlx"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	serverExitTime = time.Second
)

type Server struct {
	db     *sqlx.DB
	http   *fiber.App
	cfg    *config.Config
	logger loggerimp.Logger
	email  emailService.EmailService
}

func NewServer(cfg *config.Config, http *fiber.App, logger loggerimp.Logger, db *sqlx.DB, email emailService.EmailService) *Server {
	return &Server{
		db:     db,
		http:   http,
		cfg:    cfg,
		logger: logger,
		email:  email,
	}
}

func (s *Server) Run() error {
	go func() {
		s.logger.Infof("Server running on port %s", s.cfg.Server.RunningPort)
		s.logger.Info("Port:", s.cfg.Server.RunningPort)
		if err := s.http.Listen(":" + s.cfg.Server.RunningPort); err != nil {
			s.logger.Fatal("Error starting server:", err)
		}
	}()
	s.MapHandlers()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), serverExitTime)
	defer shutdown()

	s.logger.Info("Server exited properly")
	return s.http.ShutdownWithContext(ctx)
}
