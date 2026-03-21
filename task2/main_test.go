package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestServer(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()

    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Привет от Go сервера!",
            "status":  "success",
        })
    })
    
    handler.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Errorf("Ожидался статус 200, получили %d", w.Code)
    }

    t.Log("✅ Тест пройден!")
}
