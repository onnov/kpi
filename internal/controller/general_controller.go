package controller

import (
	"fmt"
	"net/http"
)

// set200ok формирует json ответ сервера со статусом 200 OK
func set200ok(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, msg)
}

// setError формирует json ответ ошибки
func setError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = fmt.Fprintln(w, `{"error": "`+err.Error()+`"}`)
}
