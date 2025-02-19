package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockAuthService struct {
    mock.Mock
}

func (m *MockAuthService) Authenticate(username, password string) (string, error) {
    args := m.Called(username, password)
    return args.String(0), args.Error(1)
}

func TestAuthHandler_Authenticate_Success(t *testing.T) {
    authService := new(MockAuthService)
    authService.On("Authenticate", "testuser", "password").Return("test-token", nil)

    authHandler := NewAuthHandler(authService)

    reqBody := map[string]string{
        "username": "testuser",
        "password": "password",
    }
    reqBodyBytes, _ := json.Marshal(reqBody)
    req := httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBuffer(reqBodyBytes))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()

    authHandler.Authenticate(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var response map[string]string
    json.NewDecoder(rr.Body).Decode(&response)
    assert.Equal(t, "test-token", response["token"])

    authService.AssertCalled(t, "Authenticate", "testuser", "password")
}

func TestAuthHandler_Authenticate_InvalidRequest(t *testing.T) {
    authService := new(MockAuthService)

    authHandler := NewAuthHandler(authService)

    req := httptest.NewRequest(http.MethodPost, "/api/auth", nil)
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()

    authHandler.Authenticate(rr, req)

    assert.Equal(t, http.StatusBadRequest, rr.Code)
}