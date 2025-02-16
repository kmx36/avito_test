package service

import (
    "avito_test/internal/models"
    "github.com/stretchr/testify/assert"
    "testing"
)

// TestUserService_GetUserInfo_Success - тест успешного получения информации о пользователе
func TestUserService_GetUserInfo_Success(t *testing.T) {
    // Создаем мок репозитория
    userRepo := new(MockUserRepository)

    // Настраиваем мок
    user := &models.User{ID: 1, Username: "testuser", Coins: 1000}
    userRepo.On("GetUserByID", 1).Return(user, nil)

    // Создаем сервис
    userService := NewUserService(userRepo)

    // Выполняем тест
    result, err := userService.GetUserInfo(1)
    assert.NoError(t, err) // Проверяем, что ошибки нет
    assert.Equal(t, user, result) // Проверяем, что результат соответствует ожидаемому

    // Проверяем, что мок был вызван
    userRepo.AssertCalled(t, "GetUserByID", 1)
}

// TestUserService_UpdateUserCoins_Success - тест успешного обновления баланса пользователя
func TestUserService_UpdateUserCoins_Success(t *testing.T) {
    // Создаем мок репозитория
    userRepo := new(MockUserRepository)

    // Настраиваем мок
    userRepo.On("UpdateUserCoins", 1, 500).Return(nil)

    // Создаем сервис
    userService := NewUserService(userRepo)

    // Выполняем тест
    err := userService.UpdateUserCoins(1, 500)
    assert.NoError(t, err) // Проверяем, что ошибки нет

    // Проверяем, что мок был вызван
    userRepo.AssertCalled(t, "UpdateUserCoins", 1, 500)
}