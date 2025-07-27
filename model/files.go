package model

import "time"

type FileInfo struct {
	Filename   string    `json:"filename"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type FilesResponse struct {
	Username string     `json:"username"`
	Files    []FileInfo `json:"files"`
}
