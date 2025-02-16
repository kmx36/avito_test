package repository

import (
    "database/sql"
	"avito_test/internal/models"
)

type ItemRepository struct {
    db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
    return &ItemRepository{db: db}
}

func (r *ItemRepository) GetItems() ([]models.Item, error) {
    rows, err := r.db.Query("SELECT id, name, price FROM items")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var items []models.Item
    for rows.Next() {
        var item models.Item
        if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}

func (r *ItemRepository) GetItemByName(name string) (*models.Item, error) {
    var item models.Item
    err := r.db.QueryRow("SELECT id, name, price FROM items WHERE name = $1", name).Scan(&item.ID, &item.Name, &item.Price)
    if err != nil {
        return nil, err
    }
    return &item, nil
}