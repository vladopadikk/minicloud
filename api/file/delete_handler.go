package file

import (
	"encoding/json"
	"errors"
	"minicloud/api/utils"
	"minicloud/model"
	"minicloud/service"
	"net/http"
)

func DeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, filename, err := service.GetFilename(r)
		if err != nil {
			utils.WriteJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = service.DeleteUserFile(username, filename)
		if err != nil {
			if errors.Is(err, service.ErrFileNotFound) {
				utils.WriteJSONError(w, http.StatusNotFound, err.Error())
				return
			}
			utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(model.Response{
			Username: username,
			Msg:      "Файл успешно удален",
			Status:   http.StatusOK,
		})
	}
}
