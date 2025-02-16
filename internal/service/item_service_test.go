package service

import (
    "avito_test/internal/models"
    "github.com/stretchr/testify/assert"
    "testing"
)

// TestItemService_GetItems_Success - тест успешного получения списка товаров
func TestItemService_GetItems_Success(t *testing.T) {
    // Создаем мок репозитория
    itemRepo := new(MockItemRepository)

    // Настраиваем мок
    items := []models.Item{
        {ID: 1, Name: "t-shirt", Price: 80},
        {ID: 2, Name: "cup", Price: 20},
    }
    itemRepo.On("GetItems").Return(items, nil)

    // Создаем сервис
    itemService := NewItemService(itemRepo)

    // Выполняем тест
    result, err := itemService.GetItems()
    assert.NoError(t, err) // Проверяем, что ошибки нет
    assert.Equal(t, items, result) // Проверяем, что результат соответствует ожидаемому

    // Проверяем, что мок был вызван
    itemRepo.AssertCalled(t, "GetItems")
}

// TestItemService_GetItemByName_Success - тест успешного получения товара по имени
func TestItemService_GetItemByName_Success(t *testing.T) {
    // Создаем мок репозитория
    itemRepo := new(MockItemRepository)

    // Настраиваем мок
    item := &models.Item{ID: 1, Name: "t-shirt", Price: 80}
    itemRepo.On("GetItemByName", "t-shirt").Return(item, nil)

    // Создаем сервис
    itemService := NewItemService(itemRepo)

    // Выполняем тест
    result, err := itemService.GetItemByName("t-shirt")
    assert.NoError(t, err) // Проверяем, что ошибки нет
    assert.Equal(t, item, result) // Проверяем, что результат соответствует ожидаемому

    // Проверяем, что мок был вызван
    itemRepo.AssertCalled(t, "GetItemByName", "t-shirt")
}