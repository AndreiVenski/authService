package mocks

import (
	"authService/intern/models"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) GetNewTokens(ctx context.Context, userInfo *models.UserInfo) (*models.Tokens, error) {
	args := m.Called(ctx, userInfo)
	return args.Get(0).(*models.Tokens), args.Error(1)
}

func (m *MockUseCase) RefreshAccessToken(ctx context.Context, refreshToken string, refreshTokenID uuid.UUID, ipAddr string) (string, string, error) {
	args := m.Called(ctx, refreshToken, refreshTokenID, ipAddr)
	return args.String(0), args.String(1), args.Error(2)
}
