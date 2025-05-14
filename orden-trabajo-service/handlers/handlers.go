package handlers

import (
	"fmt"
	"net/http"
)

func PublicRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Esta es una ruta pública\n")
}

func AdminRoute(w http.ResponseWriter, r *http.Request) {
	response := "Hola Mundo"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func AdminOT(w http.ResponseWriter, r *http.Request) {
	response := "Hola Mundo"
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(response))
}
