package handlers

import (
    "avito_test/internal/models"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestInfoHandler_GetUserInfo_Success(t *testing.T) {
    userService := new(MockUserService)
    itemService := new(MockItemService)
    transactionService := new(MockTransactionService)

    user := &models.User{ID: 1, Username: "testuser", Coins: 1000}
    items := []models.Item{
        {ID: 1, Name: "t-shirt", Price: 80},
        {ID: 2, Name: "cup", Price: 20},
    }
    transactions := []models.Transaction{
        {ID: 1, FromUserID: 1, ToUserID: 2, ItemID: 1, Amount: 100},
    }

    userService.On("GetUserInfo", 1).Return(user, nil)
    itemService.On("GetItems").Return(items, nil)
    transactionService.On("GetUserTransactions", 1).Return(transactions, nil)

    infoHandler := NewInfoHandler(userService, itemService, transactionService)

    req := httptest.NewRequest(http.MethodGet, "/api/info", nil)
    req = req.WithContext(WithUserID(req.Context(), 1))

    rr := httptest.NewRecorder()

    infoHandler.GetUserInfo(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    userService.AssertCalled(t, "GetUserInfo", 1)
    itemService.AssertCalled(t, "GetItems")
    transactionService.AssertCalled(t, "GetUserTransactions", 1)
}