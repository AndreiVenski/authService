package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Tokens struct {
	RefreshToken   string
	AccessToken    string
	RefreshTokenID uuid.UUID
	UserID         uuid.UUID
}

type RefreshTokenRecord struct {
	UserID         uuid.UUID
	RefreshTokenID uuid.UUID
	HashedToken    string
	Expires        time.Time
}

func (r *RefreshTokenRecord) HashToken(token string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	r.HashedToken = string(hashed)
	return nil
}

func NewRefreshTokenRecord(tokens *Tokens, expires int) *RefreshTokenRecord {
	return &RefreshTokenRecord{
		UserID:         tokens.UserID,
		RefreshTokenID: tokens.RefreshTokenID,
		Expires:        time.Now().Add(time.Duration(expires) * time.Hour),
	}
}
