package handlers

import "thainsbook/internal/models"

type Application struct {
	Users   models.UserModel
	Entries models.EntryModel
	JWT     string
}
