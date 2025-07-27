package middleware

import (
	"minicloud/storage"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Отсутствует заголовок Authorization", http.StatusUnauthorized)
			return
		}

		username, err := storage.GetUsernameByToken(authHeader)

		if err != nil {
			http.Error(w, "Неверный или истёкший токен", http.StatusUnauthorized)
			return
		}

		r.Header.Set("X-Username", username)

		next.ServeHTTP(w, r)
	}
}
