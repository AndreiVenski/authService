package server

import (
	"authService/internal/authService/delivery/http"
	"authService/internal/authService/email"
	"authService/internal/authService/repository"
	"authService/internal/authService/usecase"
	"github.com/gofiber/swagger"
)

func (s *Server) MapHandlers() {
	authRepository := repository.NewAuthRepository(s.db)

	authEmail := email.NewAuthEmail(s.email)

	authUseCase := usecase.NewAuthUseCase(s.cfg, s.logger, authRepository, authEmail)

	authHandler := http.NewAuthHandler(s.cfg, s.logger, authUseCase)

	s.http.Get("/swagger/*", swagger.HandlerDefault)

	auth := s.http.Group("/api/v1/auth")
	http.MapAuthRoutes(auth, authHandler)
}
