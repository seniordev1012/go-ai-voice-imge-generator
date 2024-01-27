package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Messages Database
func createMessagesDatabase() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", MessagesDB)
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
	//TODO: Add all queries to a file and read from that file
	_, err = db.Exec(`
		CREATE TABLE messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			sender TEXT NOT NULL,
			content TEXT DEFAULT NULL,
			audio TEXT DEFAULT NULL,
			media  VARCHAR(255) DEFAULT 'NULL',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	log.Printf("Database created successfully")

	return nil
}

// Create Keylogger Database and Table if it doesn't exist
func createKeyloggerDatabase() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", KeyboardDB)
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
	db, err := sql.Open("sqlite3", SettingsDB)
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
	_, err = db.Exec(`
		CREATE TABLE settings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			platform TEXT DEFAULT NULL,
			audioOnly INTEGER DEFAULT 1,
			theme TEXT DEFAULT 'auto',
			voice TEXT DEFAULT 'AZURESPEECH',
			language TEXT DEFAULT NULL,
			createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			accessToken TEXT DEFAULT NULL,
			refreshToken TEXT DEFAULT NULL,
			tokenExpires TIMESTAMP DEFAULT NULL   
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

// Create Sessions.db if it doesn't exist and create a table for sessions if it doesn't exist
// Token is the session token
// ID is the user ID
func createSessionsDatabase() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/sessions.db")
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
		CREATE TABLE sessions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			token TEXT NOT NULL,
			userID INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	log.Printf("Database created successfully")

	return nil
}

// Create token database and table if it doesn't exist
func createTokenDatabase() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/token.db")
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
		CREATE TABLE token (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			token TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	log.Printf("Database created successfully")

	return nil
}

// Create token database and table if it doesn't exist
// TODO Productivity database
func createProductivityDatabase() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/productivity.db")
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
		log.Println("Activity DB does not exist")
	}
	// Execute the SQL command to create the table
	_, err = db.Exec(`
		CREATE TABLE productivity (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			website TEXT NOT NULL,
			time INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	log.Printf("Activity DB created successfully")

	return nil
}

// Gallery Database
// Create gallery database and table if it doesn't exist
func createGalleryDatabase() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/gallery.db")
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
		log.Println("Gallery DB does not exist")
	}
	// Execute the SQL command to create the table
	_, err = db.Exec(`
		CREATE TABLE gallery (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			path TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	log.Printf("Gallery DB created successfully")

	return nil
}

// Mas
func createMasterMessages() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/gallery.db")
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
		log.Println("Messages DB does not exist")
	}
	// Execute the SQL command to create the table
	_, err = db.Exec(`
				CREATE TABLE messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			sender TEXT NOT NULL,
			content TEXT DEFAULT NULL,
			audio TEXT DEFAULT NULL,
			media  VARCHAR(255) DEFAULT 'NULL',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	log.Printf("Message DB created successfully")

	return nil
}

func extensionsSource() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/extensions.db")
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
		log.Println("Extensions DB does not exist")
	}
	// Execute the SQL command to create the table
	_, err = db.Exec(`
				CREATE TABLE extension (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT DEFAULT NULL,
			vendor TEXT DEFAULT NULL,
			consumer_key TEXT DEFAULT NULL,
			consumer_secret TEXT DEFAULT NULL,
			access_token TEXT DEFAULT NULL,
			access_key TEXT DEFAULT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	log.Printf("Extensions Database created successfully")

	return nil
}
