package service

import (
    "avito_test/internal/models"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
    mock.Mock
}

func (m *MockUserRepo) CreateUser(username, passwordHash string) error {
    args := m.Called(username, passwordHash)
    return args.Error(0)
}

func (m *MockUserRepo) GetUserByID(userID int) (*models.User, error) {
    args := m.Called(userID)
    return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepo) GetUserByUsername(username string) (*models.User, error) {
    args := m.Called(username)
    return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepo) UpdateUserCoins(userID, coins int) error {
    args := m.Called(userID, coins)
    return args.Error(0)
}

func TestGetUserInfo(t *testing.T) {
    mockRepo := new(MockUserRepo)
    userService := NewUserService(mockRepo)

    expectedUser := &models.User{ID: 1, Username: "testuser", Coins: 100}

    mockRepo.On("GetUserByID", 1).Return(expectedUser, nil)

    user, err := userService.GetUserInfo(1)

    assert.NoError(t, err)
    assert.Equal(t, expectedUser, user)
    mockRepo.AssertCalled(t, "GetUserByID", 1)
}

func TestUpdateUserCoins(t *testing.T) {
    mockRepo := new(MockUserRepo)
    userService := NewUserService(mockRepo)

    mockRepo.On("UpdateUserCoins", 1, 150).Return(nil)

    err := userService.UpdateUserCoins(1, 150)

    assert.NoError(t, err)
    mockRepo.AssertCalled(t, "UpdateUserCoins", 1, 150)
}