package repository

import (
    "avito_test/internal/models"
    "database/sql"
    "errors"
    "log"
)

type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) CreateUser(username, passwordHash string) error {
    _, err := r.db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, passwordHash)
    return err
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
    log.Printf("Querying user by username: username=%s", username)

    var user models.User
    err := r.db.QueryRow("SELECT id, username, password_hash, coins FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Coins)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            log.Printf("User not found: username=%s", username)
            return nil, nil
        }
        log.Printf("Error querying user by username: %v", err)
        return nil, err
    }

    log.Printf("User found: username=%s", username)

    return &user, nil
}

func (r *userRepository) GetUserByID(userID int) (*models.User, error) {
    var user models.User
    err := r.db.QueryRow("SELECT id, username, password_hash, coins FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Coins)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) UpdateUserCoins(userID, coins int) error {
    _, err := r.db.Exec("UPDATE users SET coins = $1 WHERE id = $2", coins, userID)
    return err
}