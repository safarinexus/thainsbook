package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"thainsbook/internal/utils"

	"thainsbook/internal/auth"
	"thainsbook/internal/models"

	"github.com/google/uuid"
)

func (a *Application) HandleRegister(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u models.UserRequest
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&u)
	if err != nil {
		log.Println("JSON Decode Error:", err)
		utils.WriteError(w, http.StatusBadRequest, "Unable to process request")
		return
	}

	newId := uuid.NewString()
	hashedPassword, err := auth.HashPassword(u.Password)

	err = a.Users.AddUser(newId, u.Username, hashedPassword)
	if err != nil {
		log.Println("Unable to register new user:", err)
		utils.WriteError(w, http.StatusInternalServerError, "Unable to register new user: "+u.Username)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "User " + u.Username + " registered successfully."})
}

func (a *Application) HandleLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u models.UserRequest
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&u)
	if err != nil {
		log.Println("JSON Decode Error:", err)
		utils.WriteError(w, http.StatusBadRequest, "Unable to process request")
		return
	}

	retrievedPassword, err := a.Users.GetUserPassword(u.Username)
	if err != nil {
		log.Println("Invalid Credentials.")
		utils.WriteError(w, http.StatusUnauthorized, "Invalid Credentials.")
		return
	}
	passwordHashCheck := auth.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordHashCheck {
		log.Println("Invalid Credentials.")
		utils.WriteError(w, http.StatusUnauthorized, "Invalid Credentials.")
		return
	}

	log.Println("Generating auth token...")

	tokenString, expiry, err := auth.CreateToken(u.Username, a.JWT)
	if err != nil {
		log.Println("Server Error.")
		utils.WriteError(w, http.StatusInternalServerError, "Server Error.")
		return
	}

	log.Println("Token generated. Successful User login: " + u.Username)
	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": tokenString, "expiry": expiry})
}
