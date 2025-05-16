package main

import (
	"log"
	"net/http"
	"orden-trabajo-service/adapters/persistence"
	"orden-trabajo-service/handlers"
	"orden-trabajo-service/services"
	"os"

	"orden-trabajo-service/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// Carga las variables del archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
		return
	}

	r := mux.NewRouter()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("La variable de entorno DATABASE_URL no está configurada.")
	}

	db, err := persistence.NewDBConnection(dbURL)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Crear repositorio con la conexión
	otRepo := persistence.NewPostgresRepo(db)

	otService := services.NewOTService(otRepo)     // devuelve struct
	otHandler := handlers.NewOTHandler(&otService) // pasar puntero con &

	// Middleware de autenticación.
	authMiddleware := middleware.AuthMiddleware

	// Middleware para requerir el rol de "admin".
	adminRoleMiddleware := middleware.RequireRoleMiddleware("admin")

	// Ruta pública (sin middleware).
	r.HandleFunc("/ot-service/public", handlers.PublicRoute).Methods("GET")

	// Ruta protegida que requiere autenticación y el rol de "admin".
	r.Handle("/ot-service/admin", authMiddleware(adminRoleMiddleware(http.HandlerFunc(handlers.AdminRoute)))).Methods("GET")
	r.Handle("/ot-service/admin/ot-equipos",
		authMiddleware(adminRoleMiddleware(http.HandlerFunc(otHandler.Admin_GetAllOTEquipo))),
	).Methods("GET")

	log.Println("Servidor escuchando en el puerto 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
