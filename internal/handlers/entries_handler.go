package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"thainsbook/internal/models"
	"thainsbook/internal/utils"

	"github.com/google/uuid"
)

func (a *Application) HandleGetUserEntries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := ctx.Value(UsernameKey).(string)

	entries, err := a.Entries.GetEntriesByUser(username)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error in retrieving entries.")
		return
	}

	if entries == nil {
		entries = []models.EntryResponse{}
	}

	response := map[string]interface{}{
		"entries": entries,
		"count":   len(entries),
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (a *Application) HandleCreateEntry(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var e models.EntryRequest
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		log.Println("JSON Decode Error:", err)
		utils.WriteError(w, http.StatusBadRequest, "Unable to process request.")
		return
	}

	if e.Content == "" {
		utils.WriteError(w, http.StatusBadRequest, "Cannot add empty content.")
		return
	}

	time, err := utils.ParseEntryDate(e.EntryDate)
	if err != nil {
		log.Println("Error Parsing EntryDate:", err)
		utils.WriteError(w, http.StatusInternalServerError, "Unable to process request.")
	}
	log.Println("EntryDate Processed")

	newEntry := models.Entry{
		Uid:       uuid.NewString(),
		Title:     e.Title,
		Content:   e.Content,
		EntryDate: time,
	}

	err = a.Entries.AddEntry(&newEntry)
	if err != nil {
		log.Println("Error Adding Entry:", err)
		utils.WriteError(w, http.StatusInternalServerError, "Unable to process request.")
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "New entry created successfully."})
}

func (a *Application) HandleUpdateEntry(w http.ResponseWriter, r *http.Request) {
	return
}

func (a *Application) HandleDeleteEntry(w http.ResponseWriter, r *http.Request) {
	return
}
