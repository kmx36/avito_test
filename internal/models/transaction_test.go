package models

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
)

func TestTransaction_Validate(t *testing.T) {
    tests := []struct {
        name        string
        transaction Transaction
        isValid     bool
    }{
        {
            name: "Valid transaction",
            transaction: Transaction{
                ID:         1,
                FromUserID: 1,
                ToUserID:   2,
                ItemID:     1,
                Amount:     100,
                CreatedAt:  time.Now(),
            },
            isValid: true,
        },
        {
            name: "Negative amount",
            transaction: Transaction{
                ID:         2,
                FromUserID: 1,
                ToUserID:   2,
                ItemID:     1,
                Amount:     -100,
                CreatedAt:  time.Now(),
            },
            isValid: false,
        },
        {
            name: "Zero amount",
            transaction: Transaction{
                ID:         3,
                FromUserID: 1,
                ToUserID:   2,
                ItemID:     1,
                Amount:     0,
                CreatedAt:  time.Now(),
            },
            isValid: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.isValid {
                assert.NoError(t, tt.transaction.Validate())
            } else {
                assert.Error(t, tt.transaction.Validate())
            }
        })
    }
}