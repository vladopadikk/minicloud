package file

import (
	"minicloud/api/utils"
	"minicloud/service"
	"net/http"
)

func DownloadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, filename, err := service.GetFilename(r)
		if err != nil {
			utils.WriteJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		fullPath, err := service.GetFullPathToFile(username, filename)
		if err != nil {
			utils.WriteJSONError(w, http.StatusNotFound, err.Error())
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, fullPath)
	}
}
