package service

import (
	"avito_test/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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


func TestCreateTransaction(t *testing.T) {
    mockTransactionRepo := new(MockTransactionRepository)
    mockUserRepo := new(MockUserRepository)

    transactionService := NewTransactionService(mockTransactionRepo, mockUserRepo)

    fromUserID := 1
    toUserID := 2
    itemID := 3
    amount := 100

    mockTransactionRepo.On("CreateTransaction", fromUserID, toUserID, itemID, amount).Return(nil)

    err := transactionService.CreateTransaction(fromUserID, toUserID, itemID, amount)

    assert.NoError(t, err)
    mockTransactionRepo.AssertCalled(t, "CreateTransaction", fromUserID, toUserID, itemID, amount)
}