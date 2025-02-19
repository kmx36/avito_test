package models

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
    tests := []struct {
        name    string
        user    User
        isValid bool
    }{
        {
            name: "Valid user",
            user: User{
                ID:           1,
                Username:     "testuser",
                PasswordHash: "hashedpassword",
                Coins:        1000,
            },
            isValid: true,
        },
        {
            name: "Empty username",
            user: User{
                ID:           2,
                Username:     "",
                PasswordHash: "hashedpassword",
                Coins:        1000,
            },
            isValid: false,
        },
        {
            name: "Negative coins",
            user: User{
                ID:           3,
                Username:     "testuser",
                PasswordHash: "hashedpassword",
                Coins:        -100,
            },
            isValid: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.isValid {
                assert.NoError(t, tt.user.Validate())
            } else {
                assert.Error(t, tt.user.Validate())
            }
        })
    }
}