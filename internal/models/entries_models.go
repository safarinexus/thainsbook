package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Entry struct {
	Uid       string    `json:"uid"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	EntryDate time.Time `json:"entry_date"`
}

type EntryRequest struct {
	Uid       string `json:"uid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	EntryDate string `json:"entry_date"`
}

type EntryResponse struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	EntryDate string `json:"entry_date"`
	UpdatedAt string `json:"date_updated"`
	CreatedAt string `json:"date_created"`
}

type EntryModel struct {
	DB *sql.DB
}

func (m *EntryModel) AddEntry(entry *Entry) error {
	res, err := m.DB.Exec("INSERT INTO entries (id, title, content, entry_date) VALUES (?, ?, ?, ?)", entry.Uid, entry.Title, entry.Content, entry.EntryDate)
	if err != nil {
		return fmt.Errorf("Error inserting entry: %s", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error with inserting entry: %v", id)
		return fmt.Errorf("Error inserting entry: %s", err)
	}
	return nil
}

func (m *EntryModel) GetEntriesByUser(username string) ([]EntryResponse, error) {
	var entries []EntryResponse

	rows, err := m.DB.Query("SELECT entries.title, entries.content FROM entries JOIN users ON entries.user_id = users.id WHERE users.username = ?", username)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Unable to fetch entries")
	}
	defer rows.Close()

	for rows.Next() {
		var entry EntryResponse
		if err := rows.Scan(&entry.Title, &entry.Content, &entry.EntryDate, &entry.CreatedAt, &entry.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Unable to fetch entries")
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Unable to fetch entries")
	}
	return entries, nil
}
