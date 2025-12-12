package handlers

import (
	"log"
	"net/http"
	"thainsbook/internal/utils"
)

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("Error: Resource not found: %s", r.URL.Path)

	utils.WriteError(w, http.StatusNotFound, "Resource not found.")
}

func HandleUnauthorized(w http.ResponseWriter, r *http.Request) {
	log.Printf("Error: Unauthorized: %s", r.URL.Path)

	utils.WriteError(w, http.StatusUnauthorized, "Unauthorized.")
}
