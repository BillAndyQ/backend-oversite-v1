package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// CustomClaims define la estructura de tus claims JWT.
type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var SecretKey = []byte(getSecretKey())

func getSecretKey() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET no está definida en las variables de entorno")
	}
	return []byte(secret)
}

type ContextKey string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("AuthMiddleware: Authorization header requerido")
			http.Error(w, "Authorization header requerido", http.StatusUnauthorized)
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
			}
			return SecretKey, nil
		})
		if err != nil {
			log.Printf("AuthMiddleware: Error al parsear el token: %v", err)
			http.Error(w, fmt.Sprintf("Token inválido: %v", err), http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			log.Println("AuthMiddleware: Token no es válido")
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		// Si el token es válido, extraemos la información y la ponemos en el contexto
		ctx := context.WithValue(r.Context(), ContextKey("username"), claims.Username)
		ctx = context.WithValue(ctx, ContextKey("userRole"), claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireRoleMiddleware verifica si el usuario tiene el rol especificado.
func RequireRoleMiddleware(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := r.Context().Value(ContextKey("userRole"))
			log.Printf("RequireRoleMiddleware: userRole: %v", userRole)

			if userRole == nil {
				log.Println("RequireRoleMiddleware: userRole no encontrado en el contexto")
				http.Error(w, "Prohibido", http.StatusForbidden)
				return
			}
			roleStr, ok := userRole.(string)
			if !ok {
				log.Printf("RequireRoleMiddleware: Error al convertir userRole a string: %v", userRole)
				http.Error(w, "Prohibido", http.StatusForbidden)
				return
			}
			if roleStr == requiredRole {
				next.ServeHTTP(w, r)
				return
			}
			log.Printf("RequireRoleMiddleware: Rol '%s' no coincide con el requerido '%s'", roleStr, requiredRole)
			http.Error(w, "Prohibido", http.StatusForbidden)
		})
	}
}
