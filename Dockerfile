# Используем официальный образ golang для сборки
FROM golang:1.21.1 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы проекта в контейнер
COPY . .

# Собираем проект
RUN go build -o main ./cmd/main.go

# Используем минималистичный образ для выполнения
FROM alpine:latest

# Устанавливаем рабочую директорию для финального контейнера
WORKDIR /root/

# Копируем скомпилированный бинарный файл из предыдущего контейнера
COPY --from=builder /app/main .

# Копируем конфигурационный файл из папки config
COPY config/config.yaml ./config.yaml

# Указываем команду для запуска приложения
CMD ["main"]
