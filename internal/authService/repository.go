package authService

import (
	"authService/intern/models"
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	WriteRefreshTokenRecord(ctx context.Context, refreshTokenRecord *models.RefreshTokenRecord) error
	GetRefreshTokenData(ctx context.Context, refreshTokenID uuid.UUID) (*models.RefreshTokenRecord, error)
	UpdateRefreshTokenID(ctx context.Context, refreshTokenID, newRefreshTokenID uuid.UUID) error
	GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error)
}
