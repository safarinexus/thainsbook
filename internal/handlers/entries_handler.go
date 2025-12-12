package handlers

import (
	"log"
	"net/http"

	"thainsbook/internal/utils"
)

func (a *Application) HandleGetUserEntries(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello!")

	utils.WriteJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *Application) HandleGetUserEntry(w http.ResponseWriter, r *http.Request) {
	return
}

func (a *Application) HandleCreateEntry(w http.ResponseWriter, r *http.Request) {
	return
}

func (a *Application) HandleUpdateEntry(w http.ResponseWriter, r *http.Request) {
	return
}

func (a *Application) HandleDeleteEntry(w http.ResponseWriter, r *http.Request) {
	return
}
