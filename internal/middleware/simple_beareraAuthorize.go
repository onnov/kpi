package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// SimpleBearerAuthorize простая (не для прода) авторизация по заголовку
func SimpleBearerAuthorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")
		aa := strings.Split(apiKey, " ")
		if len(aa) == 2 && aa[0] == "Bearer" && aa[1] == "48ab34464a5573519725deb5865cc74c" {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = fmt.Fprintln(w, `{"error": "Not authorized, Invalid token"}`)
		}
	})
}
