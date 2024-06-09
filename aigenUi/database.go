package aigenUi

import (
	"database/sql"
	"log"
)

const (
	MessagesDB = "DB/messages.db"
	SettingsDB = "DB/settings.db"
	KeyboardDB = "DB/keylogger.db"
	LLMDB      = "DB/llmSelection.db"
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
	_, err = db.Exec("INSERT INTO settings (selection) VALUES (?)", provider)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
	}
}

func SelectedVoiceModel() (string, error) {
	db, err := sql.Open("sqlite3", SettingsDB)
	if err != nil {
		return "", err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}

	}(db)

	var provider string
	err = db.QueryRow("SELECT selection FROM settings").Scan(&provider)
	if err != nil {
		return "", err
	}
	return provider, nil
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

func UpdateSelectedModel(selection string) error {
	db, err := sql.Open("sqlite3", LLMDB)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)
	_, err = db.Exec("UPDATE llmSelection SET selection = ? WHERE id = 1", selection)
	if err != nil {
		return err
	}
	return nil
}

func GetSelectedModel() (string, error) {
	db, err := sql.Open("sqlite3", LLMDB)
	if err != nil {
		return "", err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)
	var selection string
	err = db.QueryRow("SELECT selection FROM llmSelection").Scan(&selection)
	if err != nil {
		return "", err
	}
	return selection, nil
}
