package repository

import (
	"authService/intern/authService"
	"authService/intern/models"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
	if err != nil {
		return errors.Wrap(err, "authRepository.WriteRefreshTokenRecord.ExecContext")
	}
	return nil
}

func (r *authRepository) GetRefreshTokenData(ctx context.Context, refreshTokenID uuid.UUID) (*models.RefreshTokenRecord, error) {
	refreshTokenRecord := &models.RefreshTokenRecord{}
	var hashedToken string
	if err := r.db.QueryRowxContext(ctx, getRefreshTokenRecordQuery, refreshTokenID.String()).Scan(
		&refreshTokenRecord.UserID,
		&refreshTokenRecord.RefreshTokenID,
		&hashedToken,
		&refreshTokenRecord.Expires,
		&refreshTokenRecord.IPAddr,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(sql.ErrNoRows, "refresh token not found")
		}
		return nil, errors.Wrap(err, "authRepository.GetRefreshTokenData.Scan")
	}
	refreshTokenRecord.WriteToken(hashedToken)

	return refreshTokenRecord, nil
}

func (r *authRepository) UpdateRefreshTokenID(ctx context.Context, refreshTokenID, newRefreshTokenID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, updateRefreshTokenIDQuery, newRefreshTokenID.String(), refreshTokenID.String())
	if err != nil {
		return errors.Wrap(err, "authRepository.UpdateRefreshTokenID.ExecContext")
	}
	return nil
}

func (r *authRepository) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	user := &models.User{}
	if err := r.db.QueryRowxContext(ctx, getUserQuery, userID.String()).StructScan(user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(sql.ErrNoRows, "user not found")
		}
		return nil, errors.Wrap(err, "authRepository.GetUser.StructScan")
	}
	return user, nil
}
