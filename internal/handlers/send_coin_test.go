package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestSendCoinHandler_SendCoins_Success(t *testing.T) {
    transactionService := new(MockTransactionService)
    transactionService.On("SendCoins", 1, "user2", 100).Return(nil)

    sendCoinHandler := NewSendCoinHandler(transactionService)

    reqBody := map[string]interface{}{
        "toUser": "user2",
        "amount": 100,
    }
    reqBodyBytes, _ := json.Marshal(reqBody)
    req := httptest.NewRequest(http.MethodPost, "/api/sendCoin", bytes.NewBuffer(reqBodyBytes))
    req.Header.Set("Content-Type", "application/json")
    req = req.WithContext(WithUserID(req.Context(), 1))

    rr := httptest.NewRecorder()

    sendCoinHandler.SendCoins(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    transactionService.AssertCalled(t, "SendCoins", 1, "user2", 100)
}