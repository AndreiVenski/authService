package http

import (
	"authService/config"
	"authService/intern/authService/mocks"
	"authService/intern/models"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandler_GetNewTokens_Success(t *testing.T) {
	cfg := &config.Config{}
	mockLogger := new(mocks.MockLogger)
	mockAuthUC := new(mocks.MockUseCase)

	handler := NewAuthHandler(cfg, mockLogger, mockAuthUC)

	app := fiber.New()

	app.Post("/tokens", handler.GetNewTokens)

	userInfo := &models.UserInfo{
		UserID: uuid.New(),
	}
	requestBody, err := json.Marshal(userInfo)
	assert.NoError(t, err)

	tokens := &models.Tokens{
		AccessToken:  "mock-access-token",
		RefreshToken: "mock-refresh-token",
	}

	mockAuthUC.On("GetNewTokens", mock.Anything, mock.AnythingOfType("*models.UserInfo")).Return(tokens, nil)

	req := httptest.NewRequest(http.MethodPost, "/tokens", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Forwarded-For", "127.0.0.1")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var respBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	assert.NoError(t, err)

	tokensResp, ok := respBody["tokens"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "mock-access-token", tokensResp["access_token"])
	assert.Equal(t, "mock-refresh-token", tokensResp["refresh_token"])

	mockAuthUC.AssertExpectations(t)
}

func TestAuthHandler_RefreshTokens_Success(t *testing.T) {
	cfg := &config.Config{}
	mockLogger := new(mocks.MockLogger)
	mockAuthUC := new(mocks.MockUseCase)

	handler := NewAuthHandler(cfg, mockLogger, mockAuthUC)

	app := fiber.New()

	app.Post("/refresh", handler.RefreshTokens)

	refreshTokenBytes := make([]byte, 32)
	for i := 0; i < 32; i++ {
		refreshTokenBytes[i] = byte(i)
	}
	refreshTokenStr := base64.StdEncoding.EncodeToString(refreshTokenBytes)
	assert.Equal(t, 44, len(refreshTokenStr), "Refresh token must be 44 characters long")

	refreshTokenID := uuid.New()
	refreshTokenIDStr := refreshTokenID.String()

	refreshTokenData := map[string]interface{}{
		"refresh_token":    refreshTokenStr,
		"refresh_token_id": refreshTokenIDStr,
	}
	requestBody, err := json.Marshal(refreshTokenData)
	assert.NoError(t, err)

	newAccessToken := "new-access-token"
	newRefreshTokenID := uuid.New().String()

	mockAuthUC.On(
		"RefreshAccessToken",
		mock.Anything,
		refreshTokenStr,
		refreshTokenID,
		"127.0.0.1",
	).Return(newAccessToken, newRefreshTokenID, nil)

	req := httptest.NewRequest(http.MethodPost, "/refresh", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Forwarded-For", "127.0.0.1")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var respBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	assert.NoError(t, err)

	assert.Equal(t, newAccessToken, respBody["access_token"])
	assert.Equal(t, newRefreshTokenID, respBody["refresh_token_id"])

	mockAuthUC.AssertExpectations(t)
}

func TestAuthHandler_GetNewTokens_InvalidJSON(t *testing.T) {
	cfg := &config.Config{}
	mockLogger := new(mocks.MockLogger)
	mockAuthUC := new(mocks.MockUseCase)

	handler := NewAuthHandler(cfg, mockLogger, mockAuthUC)

	app := fiber.New()
	app.Post("/tokens", handler.GetNewTokens)

	invalidJSON := `{"invalid":`

	req := httptest.NewRequest(http.MethodPost, "/tokens", bytes.NewReader([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Forwarded-For", "127.0.0.1")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var respBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	assert.NoError(t, err)

	assert.Equal(t, "Invalid request data", respBody["error"])
}

func TestAuthHandler_GetNewTokens_NoXForwardedFor(t *testing.T) {
	cfg := &config.Config{}
	mockLogger := new(mocks.MockLogger)
	mockAuthUC := new(mocks.MockUseCase)

	handler := NewAuthHandler(cfg, mockLogger, mockAuthUC)

	app := fiber.New()
	app.Post("/tokens", handler.GetNewTokens)

	userInfo := &models.UserInfo{
		UserID: uuid.New(),
	}
	requestBody, err := json.Marshal(userInfo)
	assert.NoError(t, err)

	tokens := &models.Tokens{
		AccessToken:  "mock-access-token",
		RefreshToken: "mock-refresh-token",
	}

	mockAuthUC.On("GetNewTokens", mock.Anything, mock.AnythingOfType("*models.UserInfo")).Return(tokens, nil)

	req := httptest.NewRequest(http.MethodPost, "/tokens", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var respBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	assert.NoError(t, err)

	tokensResp, ok := respBody["tokens"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "mock-access-token", tokensResp["access_token"])
	assert.Equal(t, "mock-refresh-token", tokensResp["refresh_token"])

	mockAuthUC.AssertExpectations(t)
}

//func TestAuthHandler_GetNewTokens_AuthUCError(t *testing.T) {
//	cfg := &config.Config{}
//	mockLogger := new(mocks.MockLogger)
//	mockAuthUC := new(mocks.MockUseCase)
//
//	handler := NewAuthHandler(cfg, mockLogger, mockAuthUC)
//
//	app := fiber.New()
//	app.Post("/tokens", handler.GetNewTokens)
//
//	userInfo := &models.UserInfo{
//		UserID: uuid.New(),
//	}
//	requestBody, err := json.Marshal(userInfo)
//	assert.NoError(t, err)
//
//	mockAuthUC.On("GetNewTokens", mock.Anything, mock.AnythingOfType("*models.UserInfo")).
//		Return(nil, errors.New("mock error"))
//
//	req := httptest.NewRequest(http.MethodPost, "/tokens", bytes.NewReader(requestBody))
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("X-Forwarded-For", "127.0.0.1")
//
//	resp, err := app.Test(req)
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
//
//	var respBody map[string]interface{}
//	err = json.NewDecoder(resp.Body).Decode(&respBody)
//	assert.NoError(t, err)
//
//	assert.Equal(t, "Could not generate tokens", respBody["error"])
//
//	mockAuthUC.AssertExpectations(t)
//}

func TestAuthHandler_RefreshTokens_InvalidData(t *testing.T) {
	cfg := &config.Config{}
	mockLogger := new(mocks.MockLogger)
	mockAuthUC := new(mocks.MockUseCase)

	handler := NewAuthHandler(cfg, mockLogger, mockAuthUC)

	app := fiber.New()
	app.Post("/refresh", handler.RefreshTokens)

	invalidData := `{"refresh_token": "invalid", "refresh_token_id": "not-a-uuid"}`

	req := httptest.NewRequest(http.MethodPost, "/refresh", bytes.NewReader([]byte(invalidData)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Forwarded-For", "127.0.0.1")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var respBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	assert.NoError(t, err)

	assert.Equal(t, "Invalid refresh token data", respBody["error"])
}

func TestAuthHandler_GetNewTokens_AuthUCError(t *testing.T) {
	cfg := &config.Config{}
	mockLogger := new(mocks.MockLogger)
	mockAuthUC := new(mocks.MockUseCase)

	handler := NewAuthHandler(cfg, mockLogger, mockAuthUC)

	app := fiber.New()
	app.Post("/tokens", handler.GetNewTokens)

	userInfo := &models.UserInfo{
		UserID: uuid.New(),
	}
	requestBody, err := json.Marshal(userInfo)
	assert.NoError(t, err)

	mockAuthUC.On("GetNewTokens", mock.Anything, mock.AnythingOfType("*models.UserInfo")).
		Return((*models.Tokens)(nil), errors.New("mock error"))

	req := httptest.NewRequest(http.MethodPost, "/tokens", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Forwarded-For", "127.0.0.1")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	var respBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	assert.NoError(t, err)

	assert.Equal(t, "Could not generate tokens", respBody["error"])

	mockAuthUC.AssertExpectations(t)
}
