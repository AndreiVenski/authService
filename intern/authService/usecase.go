package authService

import "authService/intern/models"

type UseCase interface {
	GetNewTokens(userInfo *models.UserInfo) (*models.Tokens, error)
	RefreshAccessToken(refreshToken, ipAddr string) (*models.Tokens, error)
}
