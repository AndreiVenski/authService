package authService

import (
	"authService/intern/models"
	"github.com/google/uuid"
)

type Repository interface {
	WriteRefreshTokenRecord(refreshTokenRecord *models.RefreshTokenRecord) error
	GetRefreshTokenData(refreshToken string) (*models.RefreshTokenRecord, error)
	UpdateRefreshTokenRecord(refreshTokenRecord *models.RefreshTokenRecord) error
	GetUser(userID uuid.UUID) (*models.User, error)
}
