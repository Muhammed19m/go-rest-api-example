
# используем из docker hub официальный образ golang
FROM golang:1.21-alpine


# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app


# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./


# Скачиваем зависимости
RUN go mod download

# Копируем исходный код в контейнер
COPY . .

# Собираем приложение
RUN go build -o app "cmd/app/main.go"

CMD ["./app"]
