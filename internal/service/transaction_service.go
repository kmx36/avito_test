package service

import (
    "avito_test/internal/models"
    "avito_test/internal/repository"
    "errors"
)

type transactionService struct {
    transactionRepo repository.TransactionRepository
    userRepo        repository.UserRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository, userRepo repository.UserRepository) TransactionService {
    return &transactionService{transactionRepo: transactionRepo, userRepo: userRepo}
}

func (s *transactionService) SendCoins(fromUserID int, toUsername string, amount int) error {
    if amount <= 0 {
        return errors.New("amount must be positive")
    }

    fromUser, err := s.userRepo.GetUserByID(fromUserID)
    if err != nil {
        return err
    }
    if fromUser.Coins < amount {
        return errors.New("insufficient coins")
    }

    toUser, err := s.userRepo.GetUserByUsername(toUsername)
    if err != nil {
        return err
    }
    if toUser == nil {
        return errors.New("recipient not found")
    }

    if err := s.userRepo.UpdateUserCoins(fromUserID, fromUser.Coins-amount); err != nil {
        return err
    }
    if err := s.userRepo.UpdateUserCoins(toUser.ID, toUser.Coins+amount); err != nil {
        return err
    }

    return s.transactionRepo.CreateTransaction(fromUserID, toUser.ID, 0, amount)
}

func (s *transactionService) CreateTransaction(fromUserID, toUserID, itemID, amount int) error {
    return s.transactionRepo.CreateTransaction(fromUserID, toUserID, itemID, amount)
}

func (s *transactionService) GetUserTransactions(userID int) ([]models.Transaction, error) {
    return s.transactionRepo.GetUserTransactions(userID)
}