package handlers

import (
    "avito_test/internal/service"
    "net/http"

    "github.com/go-chi/chi/v5" 
)

type BuyHandler struct {
    itemService        service.ItemService
    userService        service.UserService
    transactionService service.TransactionService
}

func NewBuyHandler(itemService service.ItemService, userService service.UserService, transactionService service.TransactionService) *BuyHandler {
    return &BuyHandler{
        itemService:        itemService,
        userService:        userService,
        transactionService: transactionService,
    }
}

func (h *BuyHandler) BuyItem(w http.ResponseWriter, r *http.Request) {
    itemName := chi.URLParam(r, "item") 
    if itemName == "" {
        http.Error(w, "item name is required", http.StatusBadRequest)
        return
    }
    userID, ok := r.Context().Value("userID").(int)
    if !ok {
        http.Error(w, "userID not found in context", http.StatusInternalServerError)
        return
    }
    item, err := h.itemService.GetItemByName(itemName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if item == nil {
        http.Error(w, "item not found", http.StatusNotFound)
        return
    }
    user, err := h.userService.GetUserInfo(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if user.Coins < item.Price {
        http.Error(w, "insufficient coins", http.StatusBadRequest)
        return
    }
    newBalance := user.Coins - item.Price
    if err := h.userService.UpdateUserCoins(userID, newBalance); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := h.transactionService.CreateTransaction(userID, 0, item.ID, item.Price); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}