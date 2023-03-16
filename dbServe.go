package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Message represents a message in the database
// GetMessages retrieves all messages from the database
func getMessages() ([]Message, error) {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/messages.db")
	if err != nil {
		return nil, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}(db)

	// Execute a SQL query to retrieve all messages
	rows, err := db.Query("SELECT id, sender, content, created_at FROM messages ORDER BY created_at")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}(rows)

	// Iterate over the result set and create a slice of Message structs
	var messages []Message
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.Sender, &m.Content, &m.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, m)
		log.Printf("Message: %v", m)

	}
	return messages, nil
}
