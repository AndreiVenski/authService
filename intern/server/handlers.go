package server

import (
	"authService/intern/authService/delivery/http"
	"authService/intern/authService/email"
	"authService/intern/authService/repository"
	"authService/intern/authService/usecase"
)

func (s *Server) MapHandlers() {
	authRepository := repository.NewAuthRepository(s.db)

	authEmail := email.NewAuthEmail(s.email)

	authUseCase := usecase.NewAuthUseCase(s.cfg, s.logger, authRepository, authEmail)

	authHandler := http.NewAuthHandler(s.cfg, s.logger, authUseCase)

	auth := s.http.Group("/api/v1/auth")
	http.MapAuthRoutes(auth, authHandler)
}
