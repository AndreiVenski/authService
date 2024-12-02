package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Tokens struct {
	RefreshToken   string    `json:"refresh_token"`
	AccessToken    string    `json:"access_token"`
	RefreshTokenID uuid.UUID `json:"refresh_token_id"`
	UserID         uuid.UUID `json:"user_id"`
}

type RefreshData struct {
	Token   string    `json:"refresh_token" validate:"required,base64,len=44"`
	TokenID uuid.UUID `json:"refresh_token_id" validate:"required,uuid4"`
}

// Only for swagger used
type GetNewTokenData struct {
	UserID uuid.UUID `json:"user_id"`
}

// Only for swagger used
type TokenResponse struct {
	AccessToken    string `json:"access_token"`
	RefreshTokenID string `json:"refresh_token_id"`
}

type RefreshTokenRecord struct {
	UserID         uuid.UUID `db:"user_id"`
	RefreshTokenID uuid.UUID `db:"refresh_token_id"`
	hashedToken    string    `db:"hashed_token"`
	Expires        time.Time `db:"expires"`
	IPAddr         string    `db:"ip_addr"`
}

func (r *RefreshTokenRecord) HashToken(token string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	r.hashedToken = string(hashed)
	return nil
}

func (r *RefreshTokenRecord) WriteToken(token string) {
	r.hashedToken = token
}

func (r *RefreshTokenRecord) VerifyRefreshToken(token string) bool {
	return bcrypt.CompareHashAndPassword([]byte(r.hashedToken), []byte(token)) == nil
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
