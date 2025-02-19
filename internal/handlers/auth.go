package handlers

import (
    "net/http"
    "avito_test/internal/service"
    "encoding/json"
    "errors"
    "log"
)

type AuthHandler struct {
    authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
    return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        log.Printf("Error decoding request body: %v", err)
        http.Error(w, "invalid request body", http.StatusBadRequest)
        return
    }

    if req.Username == "" || req.Password == "" {
        log.Printf("Empty username or password")
        http.Error(w, "username and password are required", http.StatusBadRequest)
        return
    }

    token, err := h.authService.Authenticate(req.Username, req.Password)
    if err != nil {
        log.Printf("Authentication error: %v", err)
        if errors.Is(err, service.ErrUserNotFound) {
            http.Error(w, "user not found", http.StatusNotFound)
            return
        }
        if errors.Is(err, service.ErrInvalidPassword) {
            http.Error(w, "invalid password", http.StatusUnauthorized)
            return
        }
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
    })
}