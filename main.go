package main

import (
	"log"
	"minicloud/api/auth"
	"minicloud/api/file"
	"minicloud/db"
	"minicloud/middleware"
	"minicloud/service"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/register", middleware.OnlyMethod(http.MethodPost, auth.RegisterHandler()))
	mux.HandleFunc("/login", middleware.OnlyMethod(http.MethodPost, auth.LoginHandler()))
	mux.HandleFunc("/upload", middleware.OnlyMethod(http.MethodPost, middleware.AuthMiddleware(file.UploadHandler())))
	mux.HandleFunc("/download", middleware.OnlyMethod(http.MethodGet, middleware.AuthMiddleware(file.DownloadHandler())))
	mux.HandleFunc("/delete", middleware.OnlyMethod(http.MethodDelete, middleware.AuthMiddleware(file.DeleteHandler())))
	mux.HandleFunc("/files", middleware.OnlyMethod(http.MethodGet, middleware.AuthMiddleware(file.ListFilesHandler())))

	go func() {
		for {
			service.CleanExpiredSessions()
			time.Sleep(1 * time.Hour)
		}
	}()

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Ошибка запуска сервера:", err)
	}

}
