package models

import "errors"

type User struct {
    ID           int    `json:"id"`
    Username     string `json:"username"`
    PasswordHash string `json:"-"`
    Coins        int    `json:"coins"`
}

func (u *User) Validate() error {
    if u.Username == "" {
        return errors.New("username is required")
    }
    if u.Coins < 0 {
        return errors.New("coins cannot be negative")
    }
    return nil
}