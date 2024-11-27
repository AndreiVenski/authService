package repository

import (
	"authService/intern/authService"
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
