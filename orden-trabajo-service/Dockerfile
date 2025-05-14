# Etapa 1: Build del binario con Go 1.24
FROM golang:1.24-alpine AS builder

# Instala git si es necesario por dependencias
RUN apk add --no-cache git

WORKDIR /app

# Copiar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .
RUN go build -o app main.go
RUN go get github.com/joho/godotenv

# Etapa 2: Imagen mínima para producción
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .
COPY --from=builder /app/.env .

CMD ["./app"]
