package handlers

import (
	"avito_test/internal/service"
	// "encoding/json"
	"net/http"
	// "strconv"
)

type BuyHandler struct {
    itemService        *service.ItemService
    userService        *service.UserService
    transactionService *service.TransactionService
}

func NewBuyHandler(itemService *service.ItemService, userService *service.UserService, transactionService *service.TransactionService) *BuyHandler {
    return &BuyHandler{
        itemService:        itemService,
        userService:        userService,
        transactionService: transactionService,
    }
}

func (h *BuyHandler) BuyItem(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(int)
    itemName := r.PathValue("item")

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

    // Обновляем баланс пользователя
    if err := h.userService.UpdateUserCoins(userID, user.Coins-item.Price); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Создаем запись о транзакции
    if err := h.transactionService.CreateTransaction(userID, 0, item.ID, item.Price); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}