# MiniCloud

**Простое и безопасное облачное хранилище файлов на Go с аутентификацией пользователей.**

---

## Возможности

- Регистрация и аутентификация пользователей  
- Загрузка, скачивание и удаление файлов  
- Просмотр списка загруженных файлов  
- Безопасное хранение паролей (bcrypt)  
- Сессионная аутентификация с автоматической очисткой  
- Автоматическое удаление истекших сессий  

---

## Технологии

- **Язык:** Go 1.24  
- **База данных:** PostgreSQL  
- **Фреймворк:** Стандартная библиотека (`net/http`)  
- **Хеширование:** `bcrypt`  
- **UUID:** `github.com/google/uuid`

---

## Установка

### Предварительные требования

- Go 1.24 или выше  
- PostgreSQL  
- Git  

### Клонирование репозитория

```bash
git clone https://github.com/yourusername/minicloud.git
cd minicloud
```

### Установка зависимостей

```bash
go mod download
```

### Настройка базы данных

```bash
CREATE DATABASE minicloud;
```

#### Выполните SQL скрипт для создания таблиц:

```bash
psql -U postgres -d minicloud -f schema.sql
```

### Конфигурация

#### Создайте файл .env в корне проекта:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=minicloud
```

---

## Запуск
```bash
go run main.go
```

---
 
## API Документация

### Регистрация пользователя

```http
POST /register
Content-Type: application/json

{
    "username": "ivan",
    "password": "12345"
}
```

#### Ответ:

```json
{
    "username": "ivan",
    "msg": "Регистрация успешна",
    "status": 201
}
```

### Вход в систему

```http
POST /login
Content-Type: application/json

{
    "username": "ivan",
    "password": "12345"
}
```

#### Ответ:

```json
{
    "username": "ivan",
    "msg": "Логин успешен",
    "token": "uuid-token-here",
    "status": 200
}
```

### Загрузка файла

```http
POST /upload
Authorization: your-token-here
Content-Type: multipart/form-data

file: [выберите файл]
```

#### Ответ:

```json
{
    "username": "ivan",
    "msg": "Файл успешно загружен",
    "status": 201
}
```

### Скачивание файла

```http
GET /download?filename=example.txt
Authorization: your-token-here
```

### Список файлов

```http
GET /files
Authorization: your-token-here
```

#### Ответ:

```json
{
    "username": "ivan",
    "files": [
        {
            "filename": "example.txt",
            "uploaded_at": "2025-01-15T10:30:00Z"
        }
    ]
}
```

### Удаление файла

```http
DELETE /delete?filename=example.txt
Authorization: your-token-here
```

#### Ответ:

```json
{
    "username": "john_doe",
    "msg": "Файл успешно удален",
    "status": 200
}
```

## Примеры API-запросов через CURL

### Регистрация пользователя

```bash
curl -X POST http://localhost:8080/register \
     -H "Content-Type: application/json" \
     -d '{"username": "ivan", "password": "123456"}'
```

### Вход в систему

```bash
curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"username":"ivan","password":"123456"}'
```

### Загрузка файла

```bash
curl -X POST http://localhost:8080/upload \
  -H "Authorization: your-token-here" \
  -F "file=@example.txt"
```

### Скачивание файла

```bash
curl -X GET "http://localhost:8080/download?filename=example.txt" \
     -H "Authorization: your-token-here"
```

### Список файлов

```bash
curl -X GET http://localhost:8080/files \
     -H "Authorization: your-token-here"
```

### Удаление файла

```bash
curl -X DELETE "http://localhost:8080/delete?filename=example.txt" \
     -H "Authorization: your-token-here"
```
