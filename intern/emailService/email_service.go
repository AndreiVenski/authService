package emailService

import "authService/intern/models"

type EmailService interface {
	Send(email models.Email) error
}
