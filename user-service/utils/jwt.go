package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey = []byte(getSecretKey())

func getSecretKey() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}
	secret := os.Getenv("JWT_SECRET")
	log.Printf("AUTH SERVICE SECRET KEY: %s", secret)
	if secret == "" {
		panic("JWT_SECRET no está definida en las variables de entorno")
	}
	return []byte(secret)
}

func GenerateJWT(username string, role string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
