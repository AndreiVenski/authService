package authService

import "authService/intern/models"

type Email interface {
	SendWarningIPEmail(user *models.User, ipaddr string) error
}
