package handlers

import (
    "avito_test/internal/models"
    "github.com/stretchr/testify/mock"
)

type MockItemService struct {
    mock.Mock
}

func (m *MockItemService) GetItemByName(name string) (*models.Item, error) {
    args := m.Called(name)
    return args.Get(0).(*models.Item), args.Error(1)
}

func (m *MockItemService) GetItems() ([]models.Item, error) {
    args := m.Called()
    return args.Get(0).([]models.Item), args.Error(1)
}

type MockUserService struct {
    mock.Mock
}

func (m *MockUserService) GetUserInfo(userID int) (*models.User, error) {
    args := m.Called(userID)
    return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) UpdateUserCoins(userID, coins int) error {
    args := m.Called(userID, coins)
    return args.Error(0)
}

type MockTransactionService struct {
    mock.Mock
}

func (m *MockTransactionService) CreateTransaction(fromUserID, toUserID, itemID, amount int) error {
    args := m.Called(fromUserID, toUserID, itemID, amount)
    return args.Error(0)
}

func (m *MockTransactionService) GetUserTransactions(userID int) ([]models.Transaction, error) {
    args := m.Called(userID)
    return args.Get(0).([]models.Transaction), args.Error(1)
}

func (m *MockTransactionService) SendCoins(fromUserID int, toUsername string, amount int) error {
    args := m.Called(fromUserID, toUsername, amount)
    return args.Error(0)
}