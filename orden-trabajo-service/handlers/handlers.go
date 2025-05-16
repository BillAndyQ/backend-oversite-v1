package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orden-trabajo-service/services"
	"orden-trabajo-service/shared"
	"strconv"
)

type OTHandler struct {
	otService *services.OTService
}

func NewOTHandler(auth *services.OTService) *OTHandler {
	return &OTHandler{otService: auth}
}

func PublicRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Esta es una ruta pública\n")
}

func AdminRoute(w http.ResponseWriter, r *http.Request) {
	response := "Hola Mundo"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (h *OTHandler) Admin_GetAllOTEquipo(w http.ResponseWriter, r *http.Request) {
	// Leer parámetros de query (?page=1&limit=10)
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	pagination := shared.Pagination{
		Page:  page,
		Limit: limit,
	}

	lista, total, err := h.otService.ListOT_Equipo_Admin(pagination)
	if err != nil {
		http.Error(w, "Error al obtener órdenes", http.StatusInternalServerError)
		return
	}

	// Construir respuesta
	response := map[string]interface{}{
		"data":       lista,
		"total":      total,
		"page":       pagination.Page,
		"limit":      pagination.Limit,
		"totalPages": (total + pagination.Limit - 1) / pagination.Limit,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
