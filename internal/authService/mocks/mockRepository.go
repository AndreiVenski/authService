package mocks

import (
	"authService/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) WriteRefreshTokenRecord(ctx context.Context, record *models.RefreshTokenRecord) error {
	args := m.Called(ctx, record)
	return args.Error(0)
}

func (m *MockRepository) GetRefreshTokenData(ctx context.Context, tokenID uuid.UUID) (*models.RefreshTokenRecord, error) {
	args := m.Called(ctx, tokenID)
	return args.Get(0).(*models.RefreshTokenRecord), args.Error(1)
}

func (m *MockRepository) UpdateRefreshTokenID(ctx context.Context, oldID, newID uuid.UUID) error {
	args := m.Called(ctx, oldID, newID)
	return args.Error(0)
}

func (m *MockRepository) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*models.User), args.Error(1)
}
