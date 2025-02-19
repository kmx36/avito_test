package models

import (
    "errors"
    "time"
)

type Transaction struct {
    ID         int       `json:"id"`
    FromUserID int       `json:"from_user_id"`
    ToUserID   int       `json:"to_user_id"`
    ItemID     int       `json:"item_id"`
    Amount     int       `json:"amount"`
    CreatedAt  time.Time `json:"created_at"`
}

func (t *Transaction) Validate() error {
    if t.Amount <= 0 {
        return errors.New("amount must be positive")
    }
    if t.FromUserID == t.ToUserID {
        return errors.New("sender and receiver cannot be the same")
    }
    return nil
}