package email

import (
	"authService/intern/authService"
	"authService/intern/emailService"
	"authService/intern/models"
	"fmt"
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
		return err
	}
	return nil
}
