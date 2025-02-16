package service

import (
	"avito_test/internal/models"
	"avito_test/internal/repository"
)

type ItemService struct {
    itemRepo *repository.ItemRepository
}

func NewItemService(itemRepo *repository.ItemRepository) *ItemService {
    return &ItemService{itemRepo: itemRepo}
}

func (s *ItemService) GetItems() ([]models.Item, error) {
    return s.itemRepo.GetItems()
}

func (s *ItemService) GetItemByName(name string) (*models.Item, error) {
    return s.itemRepo.GetItemByName(name)
}