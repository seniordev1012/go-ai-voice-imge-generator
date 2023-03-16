package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

type Users struct {
	Id       int
	Email    string
	Password string
}

// Create the main app tabs
// The tabs are the main content of the app
// This is where All the fun happens :)
func mainApp(mapungubwe fyne.App) (*container.AppTabs, *container.Split) {
	stockMarketTab := StockTab()
	cryptoMarketTab := cTab(mapungubwe)
	forexMarketTab := marketTab(mapungubwe)
	newsTab := container.NewTabItem("News", widget.NewLabel("News Tab Content"))
	chat, aiGen := ChatTab()
	settingsTab := genSettings(mapungubwe)
	//Create the tabs container and add the tabs to it
	tabs := container.NewAppTabs(
		aiGen,
		cryptoMarketTab,
		forexMarketTab,
		newsTab,
		stockMarketTab,
		settingsTab,
	)
	//Create the bottom input box
	inputBoxContainer := bottomInputBox(chat, tabs, aiGen)
	loginBtn := widget.NewButton("Login", loginHandler())
	// Create the login form
	usernameField := widget.NewEntry()
	passwordField := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Email", Widget: usernameField},
			{Text: "Password", Widget: passwordField},
		},
		OnSubmit: loginBtn.OnTapped,
		OnCancel: func() {
			log.Println("Cancelled")
		},
	}

	loginTab := container.NewTabItem("Login", form)
	// Add the login tab to the tabs container
	tabs.Append(loginTab)

	return tabs, inputBoxContainer
}

func loginHandler() func() {
	return func() {
		email := "sifhufisg@gmail.com"
		password := "123456"

		log.Printf("Email: %s", email)
		log.Println("Password: ", password)
		db, err := dbPass()
		if err != nil {
			log.Println("Error in Connecting to the Database")
		} else {
			log.Println("Connected to the Database")
		}

		//Check if the User Exists
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

func generateToken() string {
	//Random String Generator
	//Generate a random string of length 32
	//This will be used as the session token
	b := make([]byte, 32)
	read, err := rand.Read(b)
	if err != nil {
		log.Println("Error in Generating the Session Token")
	} else {
		log.Println("Session Token Generated Successfully")
		log.Println("Bytes Read: ", read)
	}

	return base64.URLEncoding.EncodeToString(b)
}
