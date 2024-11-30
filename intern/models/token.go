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
	hashedToken    string
	Expires        time.Time
	IPAddr         string
}

func (r *RefreshTokenRecord) HashToken(token string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	r.hashedToken = string(hashed)
	return nil
}

func (r *RefreshTokenRecord) GetHashedToken() string {
	return r.hashedToken
}

func NewRefreshTokenRecord(tokens *Tokens, expires int, ipaddr string) *RefreshTokenRecord {
	return &RefreshTokenRecord{
		UserID:         tokens.UserID,
		RefreshTokenID: tokens.RefreshTokenID,
		Expires:        time.Now().Add(time.Duration(expires) * time.Hour),
		IPAddr:         ipaddr,
	}
}
