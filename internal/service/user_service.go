package service

import (
    "avito_test/internal/models"
    "avito_test/internal/repository"
)

type UserService struct {
    userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserInfo(userID int) (*models.User, error) {
    return s.userRepo.GetUserByID(userID)
}

func (s *UserService) UpdateUserCoins(userID, coins int) error {
    return s.userRepo.UpdateUserCoins(userID, coins)
}