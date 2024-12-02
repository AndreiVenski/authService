package mocks

import (
	"authService/intern/emailService"
	"authService/intern/models"
	"authService/pkg/logger"
	"sync"
)

type MockEmailService struct {
	mu     sync.Mutex
	emails map[string][]models.Email
	logger logger.Logger
}

func NewMockEmailService(logger1 logger.Logger) emailService.EmailService {
	return &MockEmailService{
		emails: make(map[string][]models.Email),
		logger: logger1,
	}
}

func (m *MockEmailService) Send(email models.Email) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.emails[email.To] = append(m.emails[email.To], email)
	m.logger.Infof("New letter for %s with text: %s", email.To, email.Body)

	return nil
}
