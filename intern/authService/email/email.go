package email

import (
	"authService/intern/authService"
	"authService/intern/models"
)

type authEmail struct{}

func NewAuthEmail() authService.Email {
	return &authEmail{}
}

func (e *authEmail) SendWarningIPEmail(user *models.User) error {
	return nil
}
