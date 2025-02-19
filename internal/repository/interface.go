package repository

import "avito_test/internal/models"

type UserRepository interface {
    CreateUser(username, passwordHash string) error
    GetUserByUsername(username string) (*models.User, error)
    GetUserByID(userID int) (*models.User, error)
    UpdateUserCoins(userID, coins int) error
}

type ItemRepository interface {
    GetItems() ([]models.Item, error)
    GetItemByName(name string) (*models.Item, error)
}

type TransactionRepository interface {
    CreateTransaction(fromUserID, toUserID, itemID, amount int) error
    GetUserTransactions(userID int) ([]models.Transaction, error)
}