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

// Get audio from database based on text content
func getAudio(content string) (string, error) {
	// Open a connection to the database
	dataSourceName := "DB/messages.db"
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return "", err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}(db)

	// Execute a SQL query to retrieve all messages
	rows, err := db.Query("SELECT audio FROM messages WHERE content = ?", content)
	if err != nil {
		return "", err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}(rows)

	// Iterate over the result set and create a slice of Message structs
	var audio string
	for rows.Next() {
		if err := rows.Scan(&audio); err != nil {
			return "", err
		}
	}
	return audio, nil
}

func kitchenLog(keylogger string) {
	//Store the keylogger in a file
	db, err := sql.Open("sqlite3", "DB/keylogger.db")
	if err != nil {
		log.Printf("Error opening database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}

	}(db)

	// Insert the keylogger into the database
	_, err = db.Exec("INSERT INTO keystrokes (textStuff) VALUES (?)", keylogger)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
	}
}

// addMessage adds a message to the database
func addMessage(sender string, content string) error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/messages.db")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}(db)

	// Prepare a SQL statement to insert the message into the database
	stmt, err := db.Prepare("INSERT INTO messages (sender, content) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Printf("Error closing statement: %v", err)
		}
	}(stmt)

	// Execute the prepared statement with the message as parameters
	_, err = stmt.Exec(sender, content)
	if err != nil {
		return err
	}

	return nil
}

func addMessageWithMedia(sender string, content string, audio string, media string) error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/messages.db")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}(db)

	// Prepare a SQL statement to insert the message into the database
	stmt, err := db.Prepare("INSERT INTO messages (sender, content, audio, media) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Printf("Error closing statement: %v", err)
		}
	}(stmt)

	// Execute the prepared statement with the message as parameters
	_, err = stmt.Exec(sender, content, audio, media)
	if err != nil {
		return err
	}

	return nil
}

// enable/disable audio
// True = enable
// False = disable
func enableAudio(key bool) {
	db, err := sql.Open("sqlite3", "DB/settings.db")
	if err != nil {
		log.Printf("Error opening database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}

	}(db)

	// Update settings in the database to enable audio
	_, err = db.Exec("UPDATE settings SET audioOnly = ? WHERE id = 1", key)
	if err != nil {
		log.Printf("Error updating database: %v", err)
	}

}

// Select last setting from database and return it for audioOnly
func getAudioSettings() bool {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/settings.db")
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return true
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}(db)
	//Select all settings from database and return the last one (should be the only one)
	rows, err := db.Query("SELECT audioOnly FROM settings ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Printf("Error selecting from database: %v", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}(rows)

	var audio bool
	for rows.Next() {
		if err := rows.Scan(&audio); err != nil {
			log.Printf("Error scanning row: %v", err)
		}
	}
	return audio

}
