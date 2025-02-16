package handlers

import (
    "net/http"
    "avito_test/internal/service"
    "encoding/json"
    //"strconv"
)

type InfoHandler struct {
    userService *service.UserService
    itemService *service.ItemService
    transactionService *service.TransactionService
}

func NewInfoHandler(userService *service.UserService, itemService *service.ItemService, transactionService *service.TransactionService) *InfoHandler {
    return &InfoHandler{
        userService:        userService,
        itemService:        itemService,
        transactionService: transactionService,
    }
}

func (h *InfoHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(int)

    user, err := h.userService.GetUserInfo(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    items, err := h.itemService.GetItems()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    transactions, err := h.transactionService.GetUserTransactions(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "coins":    user.Coins,
        "inventory": items,
        "coinHistory": map[string]interface{}{
            "received": transactions,
            "sent":    transactions,
        },
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}