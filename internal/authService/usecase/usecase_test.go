package usecase

import (
	"authService/config"
	"authService/internal/authService/mocks"
	"authService/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestAuthUseCase_RefreshAccessToken_Success(t *testing.T) {
	cfg := &config.Config{}
	mockLogger := new(mocks.MockLogger)
	mockRepo := new(mocks.MockRepository)
	mockEmail := new(mocks.MockEmail)
	uc := NewAuthUseCase(cfg, mockLogger, mockRepo, mockEmail)

	ctx := context.Background()
	refreshToken := "valid-refresh-token"
	refreshTokenID := uuid.New()
	ipAddr := "127.0.0.1"

	tokenData := &models.RefreshTokenRecord{
		UserID:  uuid.New(),
		Expires: time.Now().Add(1 * time.Hour),
		IPAddr:  ipAddr,
	}
	tokenData.HashToken(refreshToken)

	mockRepo.On("GetRefreshTokenData", ctx, refreshTokenID).Return(tokenData, nil)
	mockRepo.On("UpdateRefreshTokenID", ctx, refreshTokenID, mock.AnythingOfType("uuid.UUID")).Return(nil)

	accessToken, newRefreshTokenID, err := uc.RefreshAccessToken(ctx, refreshToken, refreshTokenID, ipAddr)
	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	assert.NotEmpty(t, newRefreshTokenID)

	mockRepo.AssertExpectations(t)
}

func TestAuthUseCase_GetNewTokens_Success(t *testing.T) {
	cfg := &config.Config{
		Server: config.ServerConfig{
			RefreshTokenExpiresHourInt: 24,
		},
	}

	mockLogger := new(mocks.MockLogger)
	mockRepo := new(mocks.MockRepository)
	mockEmail := new(mocks.MockEmail)

	uc := NewAuthUseCase(cfg, mockLogger, mockRepo, mockEmail)

	ctx := context.Background()
	userInfo := &models.UserInfo{
		UserID: uuid.New(),
		IP:     "127.0.0.1",
	}

	mockRepo.On("WriteRefreshTokenRecord", ctx, mock.AnythingOfType("*models.RefreshTokenRecord")).Return(nil)

	tokens, err := uc.GetNewTokens(ctx, userInfo)

	assert.NoError(t, err)
	assert.NotNil(t, tokens)
	assert.NotEmpty(t, tokens.AccessToken)
	assert.NotEmpty(t, tokens.RefreshToken)

	mockRepo.AssertExpectations(t)
}
