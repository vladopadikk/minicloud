package utils

import (
	"encoding/json"
	"minicloud/model"
	"net/http"
)

func WriteJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(model.ErrorResponse{
		Error:  msg,
		Status: status,
	})
}
