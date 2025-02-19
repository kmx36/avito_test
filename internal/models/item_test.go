package models

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestItem_Validate(t *testing.T) {
    tests := []struct {
        name    string
        item    Item
        isValid bool
    }{
        {
            name: "Valid item",
            item: Item{
                ID:    1,
                Name:  "t-shirt",
                Price: 80,
            },
            isValid: true,
        },
        {
            name: "Empty name",
            item: Item{
                ID:    2,
                Name:  "",
                Price: 20,
            },
            isValid: false,
        },
        {
            name: "Negative price",
            item: Item{
                ID:    3,
                Name:  "cup",
                Price: -10,
            },
            isValid: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.isValid {
                assert.NoError(t, tt.item.Validate())
            } else {
                assert.Error(t, tt.item.Validate())
            }
        })
    }
}