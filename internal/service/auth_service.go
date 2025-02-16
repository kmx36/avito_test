package service

import (
    //"avito_test/internal/models"
    "avito_test/internal/repository"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt"
    "errors"
    "log"
)

var (
    ErrUserNotFound    = errors.New("user not found")
    ErrInvalidPassword = errors.New("invalid password")
)

type AuthService struct {
    userRepo  *repository.UserRepository
    jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
    return &AuthService{userRepo: userRepo, jwtSecret: jwtSecret}
}

func (s *AuthService) Authenticate(username, password string) (string, error) {
    log.Printf("Authenticating user: username=%s", username)

    user, err := s.userRepo.GetUserByUsername(username)
    if err != nil {
        log.Printf("Error getting user by username: %v", err)
        return "", err
    }
    if user == nil {
        log.Printf("User not found: username=%s", username)
        return "", ErrUserNotFound
    }

    log.Printf("User found: username=%s, password_hash=%s", user.Username, user.PasswordHash)

    // Проверяем пароль
    err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
    if err != nil {
        log.Printf("Invalid password for user: username=%s", username)
        return "", ErrInvalidPassword
    }

    log.Printf("Password verified for user: username=%s", username)

    // Создаем JWT-токен
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
        "id":       user.ID,
    })

    signedToken, err := token.SignedString([]byte(s.jwtSecret))
    if err != nil {
        log.Printf("Error signing token: %v", err)
        return "", err
    }

    log.Printf("Token generated for user: username=%s", username)

    return signedToken, nil
}