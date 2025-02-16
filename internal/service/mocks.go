package service

import (
    "avito_test/internal/models"
    "github.com/stretchr/testify/mock"
)

// MockUserRepository - мок репозитория пользователей
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) GetUserByID(id int) (*models.User, error) {
    args := m.Called(id)
    return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*models.User, error) {
    args := m.Called(username)
    return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUserCoins(userID, coins int) error {
    args := m.Called(userID, coins)
    return args.Error(0)
}

// MockItemRepository - мок репозитория товаров
type MockItemRepository struct {
    mock.Mock
}

func (m *MockItemRepository) GetItems() ([]models.Item, error) {
    args := m.Called()
    return args.Get(0).([]models.Item), args.Error(1)
}

func (m *MockItemRepository) GetItemByName(name string) (*models.Item, error) {
    args := m.Called(name)
    return args.Get(0).(*models.Item), args.Error(1)
}

// MockTransactionRepository - мок репозитория транзакций
type MockTransactionRepository struct {
    mock.Mock
}

func (m *MockTransactionRepository) CreateTransaction(fromUserID, toUserID, itemID, amount int) error {
    args := m.Called(fromUserID, toUserID, itemID, amount)
    return args.Error(0)
}

func (m *MockTransactionRepository) GetUserTransactions(userID int) ([]models.Transaction, error) {
    args := m.Called(userID)
    return args.Get(0).([]models.Transaction), args.Error(1)
}