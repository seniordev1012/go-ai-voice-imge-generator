package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"encoding/base64"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
)

func dbPass() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	CE := CaCertPath

	rootCertPool := x509.NewCertPool()
	pem, _ := ioutil.ReadFile(CE)
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	registration := mysql.RegisterTLSConfig("custom", &tls.Config{RootCAs: rootCertPool})
	if registration != nil {
		log.Printf("Error registering TLS config: %s", registration)
	}
	var connectionString string
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&tls=custom", dbUser, dbPassword, dbHost, dbName)
	db, err := sql.Open("mysql", connectionString)

	return db, err
}

func SignInHandler(chat *fyne.Container, tabs *container.AppTabs, aiGen *container.TabItem) *container.Split {
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
	return inputBoxContainer
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
