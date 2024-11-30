package repository

import (
	"authService/intern/authService"
	"authService/intern/models"
	"context"
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

func (r *authRepository) WriteRefreshTokenRecord(ctx context.Context, refreshTokenRecord *models.RefreshTokenRecord) error {
	_, err := r.db.ExecContext(ctx, writeRefreshTokenRecordQuery, refreshTokenRecord.UserID,
		refreshTokenRecord.RefreshTokenID, refreshTokenRecord.GetHashedToken(),
		refreshTokenRecord.Expires, refreshTokenRecord.IPAddr)
	return err
}

func (r *authRepository) GetRefreshTokenData(ctx context.Context, hashedRefreshToken string) (*models.RefreshTokenRecord, error) {
	refreshTokenRecord := &models.RefreshTokenRecord{}
	if err := r.db.QueryRowxContext(ctx, getRefreshTokenRecordQuery, hashedRefreshToken).StructScan(refreshTokenRecord); err != nil {
		return nil, err
	}
	return refreshTokenRecord, nil
}

func (r *authRepository) UpdateRefreshTokenID(ctx context.Context, hashedRefreshToken string, refreshTokenID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, updateRefreshTokenIDQuery, refreshTokenID.String(), hashedRefreshToken)
	return err
}

func (r *authRepository) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	user := &models.User{}
	if err := r.db.QueryRowxContext(ctx, getUserQuery, userID.String()).StructScan(user); err != nil {
		return nil, err
	}
	return user, nil
}
