package handlers

import (
    "avito_test/internal/models"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/go-chi/chi/v5"
    "github.com/stretchr/testify/assert"
)

func TestBuyHandler_BuyItem_Success(t *testing.T) {

    itemService := new(MockItemService)
    userService := new(MockUserService)
    transactionService := new(MockTransactionService)

    item := &models.Item{ID: 1, Name: "t-shirt", Price: 80}
    user := &models.User{ID: 1, Username: "testuser", Coins: 1000}

    itemService.On("GetItemByName", "t-shirt").Return(item, nil)
    userService.On("GetUserInfo", 1).Return(user, nil)
    userService.On("UpdateUserCoins", 1, 920).Return(nil)
    transactionService.On("CreateTransaction", 1, 0, 1, 80).Return(nil)

    buyHandler := NewBuyHandler(itemService, userService, transactionService)

    req := httptest.NewRequest(http.MethodPost, "/api/buy/t-shirt", nil)
    req = req.WithContext(WithUserID(req.Context(), 1)) 

    r := chi.NewRouter()
    r.Post("/api/buy/{item}", buyHandler.BuyItem)

    rr := httptest.NewRecorder()
    r.ServeHTTP(rr, req) 

    assert.Equal(t, http.StatusOK, rr.Code)

    itemService.AssertCalled(t, "GetItemByName", "t-shirt")
    userService.AssertCalled(t, "GetUserInfo", 1)
    userService.AssertCalled(t, "UpdateUserCoins", 1, 920)
    transactionService.AssertCalled(t, "CreateTransaction", 1, 0, 1, 80)
}