package main

import (
	"database/sql"
	"log"
)

func loginHandler() func() {
	return func() {
		db, err := dbPass()
		if err != nil {
			log.Println("Error in Connecting to the Database")
		} else {
			log.Println("Connected to the Database")
		}

		//Check if the User Exists
		email := ""
		userData, err := db.Query(
			"SELECT id, email, password FROM storecustomers WHERE email = ?",
			email,
		)
		if err != nil {
			log.Println("Error in User the Database, User Does Not Exist With the Email: ", email)
		}

		var users []Users
		for userData.Next() {
			var user Users

			err = userData.Scan(&user.Id, &user.Email, &user.Password)
			if err != nil {
				log.Println("Error in Scanning the Database, User Does Not Exist With the Email: ", email)
			}
			users = append(users, user)
			log.Println(users)

		}

		//Check if the User Exists
		if len(users) == 0 {
			log.Println("User Does Not Exist With the Email: ", email)
		}
		var user User
		if userData.Next() {
			// Update the user's session token
			token := generateToken()
			//Create a new session
			db, err := sql.Open("sqlite3", "DB/sessions.db")
			if err != nil {

				log.Println("Error in Connecting to the Database")

				DBSayen := createSessionsDatabase()

				_, err := sql.Open("sqlite3", "DB/sessions.db")

				if err != nil {
					log.Println("Error in Connecting to the Database")
				}

				if DBSayen != nil {
					log.Println("Error in Creating the Database")
				}

			} else {
				log.Println("Connected to the Database")
			}
			//Create a new session
			_, err = db.Exec(
				"INSERT INTO sessions (token, user_id) VALUES (?, ?)",
				token,
				user.ID,
			)

			if err != nil {
				log.Println("Error in Creating a new Session")
			} else {
				log.Println("Session Created Successfully")
			}

			// TODO: Redirect the user to the main application view (tabs)
			//Hide login tab and show the main app tabs
		} else {
			// TODO: Display an error message to the user
		}
	}
}
