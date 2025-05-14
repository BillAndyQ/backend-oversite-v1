package http

import (
	"fmt"
	"net/http"
	"strings"
	"user-service/domain"
	"user-service/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
	secretKey   string
}

func NewAuthHandler(auth *services.AuthService, secretKey string) *AuthHandler {
	return &AuthHandler{authService: auth, secretKey: secretKey}
}

// Login manejará la autenticación y devolverá un JWT si las credenciales son correctas

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"` // Validación requerida
		Password string `json:"password" binding:"required"` // Validación requerida
	}

	// Validar datos de entrada
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Autenticar usuario y obtener token JWT
	tokenString, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Get user role
	role, err := h.authService.GetRole(req.Username)
	if err != nil {
		// Optionally continue with token only if role fetch fails
		c.JSON(http.StatusOK, gin.H{
			"ok":    true,
			"token": tokenString,
			"error": "No se pudo obtener el rol: " + err.Error(),
		})
		return
	}

	// Devolver el JWT
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"role":  string(role),
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Username   string `json:"username" binding:"required"`
		Names      string `json:"names" binding:"required"`
		Password   string `json:"password" binding:"required"`
		Dni        string `json:"dni"`
		Address    string `json:"address"`
		Phone      string `json:"phone"`
		TipoSangre string `json:"tipo_sangre"`
		Role       string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	role := domain.Role(req.Role)
	if !domain.IsValidRole(role) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rol inválido"})
		return
	}

	token, err := h.authService.Register(req.Username, req.Password, role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado correctamente",
		"token":   token,
	})
}

// Middleware para validar el token JWT
func (h *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Falta el token de autorización"})
			c.Abort()
			return
		}

		// El token debe venir en el formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
			}
			return []byte(h.secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Extraer el username de los claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims inválidos"})
			c.Abort()
			return
		}

		username, ok := claims["username"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username no encontrado en el token"})
			c.Abort()
			return
		}

		// Guardar el username en el contexto para usarlo en los manejadores
		c.Set("username", username)
		c.Next()
	}
}

// Ruta GET /me
func (h *AuthHandler) Me(c *gin.Context) {
	// Obtener el username del contexto (establecido por el middleware)
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	// Obtener los datos del usuario desde el servicio
	user, err := h.authService.GetUserByUsername(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Crear la respuesta, excluyendo la contraseña
	response := gin.H{
		"username":    user.Username,
		"role":        user.Role,
		"names":       user.Names,
		"dni":         user.Dni,
		"address":     user.Address,
		"phone":       user.Phone,
		"tipo_sangre": user.TipoSangre,
	}

	c.JSON(http.StatusOK, response)
}
