package email

import (
	"authService/internal/authService"
	"authService/internal/emailService"
	"authService/internal/models"
	"fmt"
	"github.com/pkg/errors"
)

type authEmail struct {
	emailService emailService.EmailService
}

func NewAuthEmail(emailService emailService.EmailService) authService.Email {
	return &authEmail{
		emailService: emailService,
	}
}

func (e *authEmail) SendWarningIPEmail(user *models.User, ipAddr string) error {
	email := models.Email{
		Body: fmt.Sprintf("Dear,%s\nSomebody with ip : %s used your refresh token. If it's not you...", user.Name, ipAddr),
		To:   user.Email,
	}

	err := e.emailService.Send(email)
	if err != nil {
		return errors.Wrap(err, "authEmail.SendWarningIPEmail.Send")
	}
	return nil
}
