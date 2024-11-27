package utils

import (
	"authService/intern/models"
	"crypto/rand"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenClaims struct {
	UserID         uuid.UUID
	IP             string
	refreshTokenID string
	jwt.RegisteredClaims
}

func generateRefreshTokenID() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(tokenBytes), nil
}

func GenerateRefreshToken(userInfo *models.UserInfo) (string, error) {
	return "", nil
}

func GenerateAccessToken(userInfo *models.UserInfo) (string, error) {

	return "", nil
}
