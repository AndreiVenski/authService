package emailService

import "authService/internal/models"

type EmailService interface {
	Send(email models.Email) error
}
