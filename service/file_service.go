package service

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"minicloud/model"
	"minicloud/storage"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/lib/pq"
)

var ErrFileExists = errors.New("file already exists")
var ErrFileNotFound = errors.New("file not found")

func SaveUserFile(username, filename string, file io.Reader) error {
	user, err := storage.GetUserByUsername(username)
	if err != nil {
		return err
	}

	dir := filepath.Join("data", username)
	fullPath := filepath.Join(".", "data", username, filename)

	err = storage.SaveFileToDisk(dir, fullPath, file)

	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	err = storage.InsertFileRecord(user.ID, filename, fullPath)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrFileExists
		}
		_ = storage.DeleteFile(fullPath)
		return fmt.Errorf("failed to insert file record: %w", err)
	}

	return nil
}

func DeleteUserFile(username, filename string) error {
	user, err := storage.GetUserByUsername(username)
	if err != nil {
		return err
	}

	fullPath, err := GetFullPathToFile(username, filename)
	if err != nil {
		return err
	}

	err = storage.DeleteFileRecord(user.ID, filename)
	if err != nil {
		return fmt.Errorf("failed to delete file record from database: %w", err)
	}

	err = storage.DeleteFile(fullPath)
	if err != nil {
		_ = storage.InsertFileRecord(user.ID, filename, fullPath)
		return errors.New("failed to delete file from disk")
	}

	return nil
}

func ListUserFiles(username string) ([]model.FileInfo, error) {
	user, err := storage.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	userFiles, err := storage.GetUserFiles(user.ID)
	if err != nil {
		return nil, errors.New("failed to get files")
	}

	return userFiles, nil
}

func GetFilename(r *http.Request) (string, string, error) {
	username := r.Header.Get("X-Username")
	filename := r.URL.Query().Get("filename")

	if filename == "" {
		return "", "", errors.New("file name is not specified")
	}
	if strings.Contains(filename, "..") {
		return "", "", errors.New("invalid file name")
	}

	return username, filename, nil
}

func GetFullPathToFile(username, filename string) (string, error) {
	storedName, err := storage.GetStoredName(filename, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrFileNotFound
		}
		return "", fmt.Errorf("failed to get stored name: %w", err)
	}

	fullPath := filepath.Join(".", storedName)
	return fullPath, nil
}
