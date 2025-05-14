package main

import (
	"log"
	"os"
	"user-service/adapters/http"
	"user-service/adapters/persistence"
	"user-service/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando .env")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("La variable de entorno DATABASE_URL no está configurada.")
	}

	// Establecer la conexión a la base de datos
	db, err := persistence.NewDBConnection(dbURL)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close() // Cierra la conexión obtenida de NewDBConnection

	// Crear el repositorio PostgreSQL pasando la conexión
	repo := persistence.NewPostgresUserRepo(db)

	// Crear el servicio de autenticación
	authService := services.NewAuthService(repo)

	// Crear el handler HTTP
	authHandler := http.NewAuthHandler(authService, os.Getenv("JWT_SECRET"))

	// Crear el router de Gin
	router := gin.Default()

	// Definir las rutas
	router.POST("/user/login", authHandler.Login)
	router.POST("/user/register", authHandler.Register)

	protected := router.Group("/").Use(authHandler.AuthMiddleware())
	protected.GET("/user/me", authHandler.Me)
	// Iniciar el servidor
	log.Println("Servidor corriendo en el puerto 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("No se pudo iniciar el servidor: ", err)
	}
}
