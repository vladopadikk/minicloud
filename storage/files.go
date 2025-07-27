package storage

import (
	"io"
	"minicloud/db"
	"minicloud/model"
	"os"
	"time"
)

func SaveFileToDisk(dir, path string, file io.Reader) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}

func InsertFileRecord(userID int, filename, storedName string) error {
	_, err := db.DB.Exec(`
		INSERT INTO files (user_id, filename, stored_name, uploaded_at) 
		VALUES ($1, $2, $3, $4)
	`, userID, filename, storedName, time.Now())
	return err
}

func GetStoredName(filename, username string) (string, error) {
	var storedName string

	err := db.DB.QueryRow(`
		SELECT f.stored_name 
		FROM files f
		JOIN users u ON u.id = f.user_id
		WHERE f.filename = $1 AND u.username = $2 
	`, filename, username).Scan(&storedName)

	if err != nil {
		return "", err
	}
	return storedName, nil
}

func DeleteFileRecord(userID int, filename string) error {
	_, err := db.DB.Exec(`
		DELETE FROM files
		WHERE user_id = $1 AND filename = $2
	`, userID, filename)
	return err
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

func GetUserFiles(userID int) ([]model.FileInfo, error) {
	rows, err := db.DB.Query(`
		SELECT f.filename, f.uploaded_at
		FROM files f
		WHERE f.user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []model.FileInfo
	for rows.Next() {
		var file model.FileInfo
		if err := rows.Scan(&file.Filename, &file.UploadedAt); err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}
