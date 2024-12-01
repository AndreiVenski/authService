package mocks

import (
	"authService/intern/models"
	"github.com/stretchr/testify/mock"
)

type MockEmail struct {
	mock.Mock
}

func (m *MockEmail) SendWarningIPEmail(user *models.User, ipAddr string) error {
	args := m.Called(user, ipAddr)
	return args.Error(0)
}
