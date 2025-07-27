package auth

import (
	"encoding/json"
	"minicloud/api/utils"
	"minicloud/model"
	"minicloud/service"
	"net/http"
)

func RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		err := json.NewDecoder(r.Body).Decode(&u)
		defer r.Body.Close()

		if err != nil {
			utils.WriteJSONError(w, http.StatusBadRequest, "Невалидный JSON")
			return
		}

		err = service.Register(u)
		if err != nil {
			utils.WriteJSONError(w, http.StatusConflict, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(model.Response{
			Username: u.Username,
			Msg:      "Регистрация успешна",
			Status:   http.StatusCreated,
		})
	}
}

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.User
		err := json.NewDecoder(r.Body).Decode(&req)
		defer r.Body.Close()

		if err != nil {
			utils.WriteJSONError(w, http.StatusBadRequest, "Невалидный JSON")
			return
		}

		token, user, err := service.Login(req)
		if err != nil {
			utils.WriteJSONError(w, http.StatusUnauthorized, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(model.Response{
			Username: user.Username,
			Msg:      "Логин успешен",
			Token:    token,
			Status:   http.StatusOK,
		})
	}
}
