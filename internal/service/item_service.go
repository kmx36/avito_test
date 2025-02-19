package service

import (
    "avito_test/internal/models"
    "avito_test/internal/repository"
)

type itemService struct {
    itemRepo repository.ItemRepository
}

func NewItemService(itemRepo repository.ItemRepository) ItemService {
    return &itemService{itemRepo: itemRepo}
}

func (s *itemService) GetItemByName(name string) (*models.Item, error) {
    return s.itemRepo.GetItemByName(name)
}

func (s *itemService) GetItems() ([]models.Item, error) {
    return s.itemRepo.GetItems()
}