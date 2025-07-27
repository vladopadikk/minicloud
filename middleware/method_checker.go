package middleware

import "net/http"

func OnlyMethod(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			http.Error(w, "Метод должен быть "+method, http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}
