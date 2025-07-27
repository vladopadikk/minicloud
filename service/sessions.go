package service

import (
	"log"
	"minicloud/db"
)

func CleanExpiredSessions() {
	_, err := db.DB.Exec(`
		DELETE FROM sessions 
		WHERE expires_at < NOW()
	`)
	if err != nil {
		log.Println("Ошибка при очистке просроченных сессий:", err)
	} else {
		log.Println("Просроченные сессии удалены")
	}
}
