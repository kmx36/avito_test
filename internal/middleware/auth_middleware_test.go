package middleware

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/golang-jwt/jwt"
    "github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_Success(t *testing.T) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": 1,
    })
    tokenString, _ := token.SignedString([]byte("your_jwt_secret"))

    req := httptest.NewRequest(http.MethodGet, "/api/info", nil)
    req.Header.Set("Authorization", "Bearer "+tokenString)

    rr := httptest.NewRecorder()

    nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(int)
        assert.Equal(t, 1, userID)
        w.WriteHeader(http.StatusOK)
    })

    middleware := AuthMiddleware("your_jwt_secret")(nextHandler)
    middleware.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/api/info", nil)
    req.Header.Set("Authorization", "Bearer invalid-token")

    rr := httptest.NewRecorder()

    nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        t.Fatal("Handler should not be called")
    })

    middleware := AuthMiddleware("your_jwt_secret")(nextHandler)
    middleware.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusUnauthorized, rr.Code)
}