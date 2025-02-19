package models

import "errors"

type Item struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}

func (i *Item) Validate() error {
    if i.Name == "" {
        return errors.New("name is required")
    }
    if i.Price <= 0 {
        return errors.New("price must be positive")
    }
    return nil
}