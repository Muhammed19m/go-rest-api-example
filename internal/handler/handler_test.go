package handler

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestHandleTransaction(t *testing.T) {

// 	payload := []byte(`{"walletId":123, "operationType":"DEPOSIT", "amount":1000}`)
// 	req, err := http.NewRequest("POST", "/api/v1/wallet", bytes.NewBuffer(payload))
// 	if err != nil {
// 		t.Fatalf("Failed to create request: %v", err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(HandleTransaction)
// 	handler.ServeHTTP(rr, req)

// }

// func TestGetBalance(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/api/v1/wallets{1}")
// 	if err != nil {
// 		t.Fatalf("Failed to create request: %v", err)
// 	}

// }
 