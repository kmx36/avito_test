package service

import (
    "avito_test/internal/models"
    "github.com/stretchr/testify/assert"
    "testing"
)

// TestTransactionService_SendCoins_Success - тест успешной передачи монет
func TestTransactionService_SendCoins_Success(t *testing.T) {
    // Создаем моки репозиториев
    userRepo := new(MockUserRepository)
    transactionRepo := new(MockTransactionRepository)

    // Настраиваем моки
    fromUser := &models.User{ID: 1, Username: "user1", Coins: 1000}
    toUser := &models.User{ID: 2, Username: "user2", Coins: 1000}
    userRepo.On("GetUserByID", 1).Return(fromUser, nil)
    userRepo.On("GetUserByUsername", "user2").Return(toUser, nil)
    userRepo.On("UpdateUserCoins", 1, 900).Return(nil)
    userRepo.On("UpdateUserCoins", 2, 1100).Return(nil)
    transactionRepo.On("CreateTransaction", 1, 2, 0, 100).Return(nil)

    // Создаем сервис
    transactionService := NewTransactionService(transactionRepo, userRepo)

    // Выполняем тест
    err := transactionService.SendCoins(1, "user2", 100)
    assert.NoError(t, err) // Проверяем, что ошибки нет

    // Проверяем, что моки были вызваны
    userRepo.AssertCalled(t, "GetUserByID", 1)
    userRepo.AssertCalled(t, "GetUserByUsername", "user2")
    userRepo.AssertCalled(t, "UpdateUserCoins", 1, 900)
    userRepo.AssertCalled(t, "UpdateUserCoins", 2, 1100)
    transactionRepo.AssertCalled(t, "CreateTransaction", 1, 2, 0, 100)
}