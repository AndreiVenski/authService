package repository

import (
	"authService/intern/authService"
	"authService/intern/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) authService.Repository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) WriteRefreshTokenRecord(refreshTokenRecord *models.RefreshTokenRecord) error {
	return nil
}
func (r *authRepository) GetRefreshTokenData(refreshToken string) (*models.RefreshTokenRecord, error) {
	return nil, nil
}
func (r *authRepository) UpdateRefreshTokenRecord(refreshTokenRecord *models.RefreshTokenRecord) error {
	return nil
}
func (r *authRepository) GetUser(userID uuid.UUID) (*models.User, error) {
	return nil, nil
}
