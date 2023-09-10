package aigenAudioAutoPlay

import (
	"database/sql"
	"log"
)

const (
	MessagesDB = "DB/messages.db"
	SettingsDB = "DB/settings.db"
	KeyboardDB = "DB/keylogger.db"
)

// UpdateBotChatAudioPath updates the audio path in the messages table
// for the last row in the table
func UpdateBotChatAudioPath(audioPath string) (string, error) {
	// SQL update audio for the last row in the messages table
	db, err := sql.Open("sqlite3", MessagesDB)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return "", err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}(db)

	_, err = db.Exec("UPDATE messages SET audio = ? WHERE id = (SELECT id FROM messages ORDER BY id DESC LIMIT 1)", audioPath)
	if err != nil {
		return "", err
	}
	return audioPath, nil
}
