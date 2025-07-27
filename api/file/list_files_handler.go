package file

import (
	"encoding/json"
	"minicloud/api/utils"
	"minicloud/model"
	"minicloud/service"
	"net/http"
)

func ListFilesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("X-Username")

		files, err := service.ListUserFiles(username)
		if err != nil {
			utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(model.FilesResponse{
			Username: username,
			Files:    files,
		})
	}
}
