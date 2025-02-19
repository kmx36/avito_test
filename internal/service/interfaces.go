package service

import "avito_test/internal/models"

type ItemService interface {
    GetItemByName(name string) (*models.Item, error)
    GetItems() ([]models.Item, error)
}

type UserService interface {
    GetUserInfo(userID int) (*models.User, error)
    UpdateUserCoins(userID, coins int) error
}

type TransactionService interface {
    CreateTransaction(fromUserID, toUserID, itemID, amount int) error
    GetUserTransactions(userID int) ([]models.Transaction, error)
	SendCoins(fromUserID int, toUsername string, amount int) error
}