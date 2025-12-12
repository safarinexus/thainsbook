package models

import "database/sql"

type Entry struct {
	Uid       string `json:"uid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"date_created"`
	UpdatedAt string `json:"date_updated"`
}

type EntryRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type EntryModel struct {
	DB *sql.DB
}
