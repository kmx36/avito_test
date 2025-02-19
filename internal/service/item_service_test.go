package service

import (
	"avito_test/internal/models"
	//"avito_test/internal/repository"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func TestGetItems(t *testing.T) {
	mockRepo := new(MockItemRepository)
	itemService := NewItemService(mockRepo)

	expectedItems := []models.Item{
		{ID: 1, Name: "T-shirt", Price: 100},
		{ID: 2, Name: "Mug", Price: 50},
	}

	mockRepo.On("GetItems").Return(expectedItems, nil)

	items, err := itemService.GetItems()

	assert.NoError(t, err)
	assert.Equal(t, expectedItems, items)
}

func TestGetItemByName(t *testing.T) {
	mockRepo := new(MockItemRepository)
	itemService := NewItemService(mockRepo)

	expectedItem := &models.Item{ID: 1, Name: "T-shirt", Price: 100}

	mockRepo.On("GetItemByName", "T-shirt").Return(expectedItem, nil)

	item, err := itemService.GetItemByName("T-shirt")

	assert.NoError(t, err)
	assert.Equal(t, expectedItem, item)
}
