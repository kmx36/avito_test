package handlers

import (
    "net/http"
    "avito_test/internal/service"
    "encoding/json"
)

type SendCoinHandler struct {
    transactionService *service.TransactionService
}

func NewSendCoinHandler(transactionService *service.TransactionService) *SendCoinHandler {
    return &SendCoinHandler{transactionService: transactionService}
}

func (h *SendCoinHandler) SendCoins(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(int)

    var req struct {
        ToUser string `json:"toUser"`
        Amount int    `json:"amount"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    if err := h.transactionService.SendCoins(userID, req.ToUser, req.Amount); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusOK)
}