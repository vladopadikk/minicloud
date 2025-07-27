package file

import (
	"encoding/json"
	"errors"
	"minicloud/api/utils"
	"minicloud/model"
	"minicloud/service"
	"net/http"
)

func UploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("X-Username")

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			utils.WriteJSONError(w, http.StatusBadRequest, "Файл слишком большой")
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			utils.WriteJSONError(w, http.StatusBadRequest, "Не удалось получить файл")
			return
		}
		defer file.Close()

		err = service.SaveUserFile(username, header.Filename, file)
		if err != nil {
			if errors.Is(err, service.ErrFileExists) {
				utils.WriteJSONError(w, http.StatusConflict, err.Error())
				return
			}
			utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(model.Response{
			Username: username,
			Msg:      "Файл успешно загружен",
			Status:   http.StatusCreated,
		})
	}
}
