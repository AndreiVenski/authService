package authService

import "authService/internal/models"

type Email interface {
	SendWarningIPEmail(user *models.User, ipaddr string) error
}
