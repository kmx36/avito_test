package service

import (
    "avito_test/internal/models"
    "github.com/stretchr/testify/assert"
    "golang.org/x/crypto/bcrypt"
    "testing"
)

// TestAuthService_Authenticate_Success - тест успешной аутентификации
func TestAuthService_Authenticate_Success(t *testing.T) {
    // Создаем мок репозитория
    userRepo := new(MockUserRepository)

    // Настраиваем мок
    password := "password"
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    userRepo.On("GetUserByUsername", "testuser").Return(&models.User{
        ID:           1,
        Username:     "testuser",
        PasswordHash: string(hashedPassword),
    }, nil)

    // Создаем сервис
    authService := NewAuthService(userRepo, "secret")

    // Выполняем тест
    token, err := authService.Authenticate("testuser", "password")
    assert.NoError(t, err) // Проверяем, что ошибки нет
    assert.NotEmpty(t, token) // Проверяем, что токен не пустой

    // Проверяем, что мок был вызван
    userRepo.AssertCalled(t, "GetUserByUsername", "testuser")
}

// TestAuthService_Authenticate_UserNotFound - тест, когда пользователь не найден
func TestAuthService_Authenticate_UserNotFound(t *testing.T) {
    // Создаем мок репозитория
    userRepo := new(MockUserRepository)

    // Настраиваем мок
    userRepo.On("GetUserByUsername", "unknownuser").Return((*models.User)(nil), nil)

    // Создаем сервис
    authService := NewAuthService(userRepo, "secret")

    // Выполняем тест
    token, err := authService.Authenticate("unknownuser", "password")
    assert.Error(t, err) // Проверяем, что ошибка есть
    assert.Equal(t, ErrUserNotFound, err) // Проверяем, что ошибка соответствует ожидаемой
    assert.Empty(t, token) // Проверяем, что токен пустой

    // Проверяем, что мок был вызван
    userRepo.AssertCalled(t, "GetUserByUsername", "unknownuser")
}

// TestAuthService_Authenticate_InvalidPassword - тест, когда пароль неверный
func TestAuthService_Authenticate_InvalidPassword(t *testing.T) {
    // Создаем мок репозитория
    userRepo := new(MockUserRepository)

    // Настраиваем мок
    password := "password"
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    userRepo.On("GetUserByUsername", "testuser").Return(&models.User{
        ID:           1,
        Username:     "testuser",
        PasswordHash: string(hashedPassword),
    }, nil)

    // Создаем сервис
    authService := NewAuthService(userRepo, "secret")

    // Выполняем тест
    token, err := authService.Authenticate("testuser", "wrongpassword")
    assert.Error(t, err) // Проверяем, что ошибка есть
    assert.Equal(t, ErrInvalidPassword, err) // Проверяем, что ошибка соответствует ожидаемой
    assert.Empty(t, token) // Проверяем, что токен пустой

    // Проверяем, что мок был вызван
    userRepo.AssertCalled(t, "GetUserByUsername", "testuser")
}