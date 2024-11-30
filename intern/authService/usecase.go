package authService

import (
	"authService/intern/models"
	"context"
)

type UseCase interface {
	GetNewTokens(ctx context.Context, userInfo *models.UserInfo) (*models.Tokens, error)
	RefreshAccessToken(ctx context.Context, refreshToken, ipAddr string) (string, error)
}
