package main

import (
	"log"
	"net/http"
	"orden-trabajo-service/handlers"

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

	// Middleware de autenticación.
	authMiddleware := middleware.AuthMiddleware

	// Middleware para requerir el rol de "admin".
	adminRoleMiddleware := middleware.RequireRoleMiddleware("administrador")

	// Ruta pública (sin middleware).
	r.HandleFunc("/ot-service/public", handlers.PublicRoute).Methods("GET")

	// Ruta protegida que requiere autenticación y el rol de "admin".
	r.Handle("/ot-service/admin", authMiddleware(adminRoleMiddleware(http.HandlerFunc(handlers.AdminRoute)))).Methods("GET")

	log.Println("Servidor escuchando en el puerto 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
