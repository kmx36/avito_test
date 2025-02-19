package service

import (
    "avito_test/internal/models"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) CreateUser(username, passwordHash string) error {
    args := m.Called(username, passwordHash)
    return args.Error(0)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*models.User, error) {
    args := m.Called(username)
    return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(userID int) (*models.User, error) {
    args := m.Called(userID)
    return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUserCoins(userID, coins int) error {
    args := m.Called(userID, coins)
    return args.Error(0)
}

func TestAuthenticate_Success(t *testing.T) {
    mockRepo := new(MockUserRepository)
    authService := NewAuthService(mockRepo, "secret")

    user := &models.User{
        ID:           1,
        Username:     "testuser",
        PasswordHash: "$2a$10$P7z6JdrxEZjDOf6V55rcHeLbzIuY6K97f5c/xqEzDYYBygws5zmmC",
    }

    mockRepo.On("GetUserByUsername", "testuser").Return(user, nil)

    token, err := authService.Authenticate("testuser", "password")

    assert.NoError(t, err)
    assert.NotEmpty(t, token)
}

func TestAuthenticate_InvalidPassword(t *testing.T) {
    mockRepo := new(MockUserRepository)
    authService := NewAuthService(mockRepo, "secret")

    user := &models.User{
        ID:           1,
        Username:     "testuser",
        PasswordHash: "$2a$10$P7z6JdrxEZjDOf6V55rcHeLbzIuY6K97f5c/xqEzDYYBygws5zmmC",
    }

    mockRepo.On("GetUserByUsername", "testuser").Return(user, nil)

    token, err := authService.Authenticate("testuser", "wrongpassword")

    assert.Error(t, err)
    assert.Empty(t, token)
    assert.Equal(t, ErrInvalidPassword, err)
}

func TestAuthenticate_UserNotFound(t *testing.T) {
    mockRepo := new(MockUserRepository)
    authService := NewAuthService(mockRepo, "secret")

    mockRepo.On("GetUserByUsername", "nonexistentuser").Return(nil, nil)

    token, err := authService.Authenticate("nonexistentuser", "password")

    assert.Error(t, err)
    assert.Empty(t, token)
    assert.Equal(t, ErrUserNotFound, err)
}