package utils

import (
	"authService/config"
	"authService/intern/models"
	"crypto/rand"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const (
	expTime = time.Hour
)

type TokenClaims struct {
	UserID         uuid.UUID
	refreshTokenID uuid.UUID
	IP             string
	jwt.RegisteredClaims
}

func GenerateTokens(cfg *config.Config, userInfo *models.UserInfo) (*models.Tokens, error) {
	refreshTokenID := uuid.New()
	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	accessToken, err := GenerateAccessToken(cfg, userInfo, refreshTokenID)
	if err != nil {
		return nil, err
	}

	return &models.Tokens{
		RefreshToken:   refreshToken,
		AccessToken:    accessToken,
		RefreshTokenID: refreshTokenID,
		UserID:         userInfo.UserID,
	}, nil
}

func GenerateRefreshToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(tokenBytes), nil
}

func GenerateAccessToken(cfg *config.Config, userInfo *models.UserInfo, refreshTokenID uuid.UUID) (string, error) {
	claims := TokenClaims{
		UserID:         userInfo.UserID,
		IP:             userInfo.IP,
		refreshTokenID: refreshTokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.Server.AccessTokenExpiresHourInt) * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(cfg.Server.JWTSecret))

	return tokenString, err
}
