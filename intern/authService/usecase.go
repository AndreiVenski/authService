package authService

import (
	"authService/intern/models"
	"context"
	"github.com/google/uuid"
)

type UseCase interface {
	GetNewTokens(ctx context.Context, userInfo *models.UserInfo) (*models.Tokens, error)
	RefreshAccessToken(ctx context.Context, refreshToken string, refreshTokenID uuid.UUID, ipAddr string) (string, string, error)
}
