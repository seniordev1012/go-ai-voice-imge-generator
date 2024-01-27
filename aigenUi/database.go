package aigenUi

import (
	"database/sql"
	"log"
)

const (
	MessagesDB = "DB/messages.db"
	SettingsDB = "DB/settings.db"
	KeyboardDB = "DB/keylogger.db"
)

func ChangeSetting(soundIsOn int) {
	db, err := sql.Open("sqlite3", SettingsDB)
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
	_, err = db.Exec("INSERT INTO settings (audioOnly) VALUES (?)", soundIsOn)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
	}
}

func ChangeVoice(provider string) {
	db, err := sql.Open("sqlite3", SettingsDB)
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
	_, err = db.Exec("INSERT INTO settings (speech_provider) VALUES (?)", provider)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
	}
}

func SoundIsOffON(soundIsOff int) {
	db, err := sql.Open("sqlite3", SettingsDB)
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
	_, err = db.Exec("INSERT INTO settings (audioOnly) VALUES (?)", soundIsOff)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
	}
}
