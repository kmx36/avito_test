package repository

import (
    "avito_test/internal/models"
    "database/sql"
    "time"
)

type transactionRepository struct {
    db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
    return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(fromUserID, toUserID, itemID, amount int) error {
    _, err := r.db.Exec("INSERT INTO transactions (from_user_id, to_user_id, item_id, amount, created_at) VALUES ($1, $2, $3, $4, $5)",
        fromUserID, toUserID, itemID, amount, time.Now())
    return err
}

func (r *transactionRepository) GetUserTransactions(userID int) ([]models.Transaction, error) {
    rows, err := r.db.Query("SELECT id, from_user_id, to_user_id, item_id, amount, created_at FROM transactions WHERE from_user_id = $1 OR to_user_id = $1", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var transactions []models.Transaction
    for rows.Next() {
        var transaction models.Transaction
        if err := rows.Scan(&transaction.ID, &transaction.FromUserID, &transaction.ToUserID, &transaction.ItemID, &transaction.Amount, &transaction.CreatedAt); err != nil {
            return nil, err
        }
        transactions = append(transactions, transaction)
    }
    return transactions, nil
}