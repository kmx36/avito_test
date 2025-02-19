package service

import (
    "avito_test/internal/models"
    "avito_test/internal/repository"
)

type userService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
    return &userService{userRepo: userRepo}
}

func (s *userService) GetUserInfo(userID int) (*models.User, error) {
    return s.userRepo.GetUserByID(userID)
}

func (s *userService) UpdateUserCoins(userID, coins int) error {
    return s.userRepo.UpdateUserCoins(userID, coins)
}