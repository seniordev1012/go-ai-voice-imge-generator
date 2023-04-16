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
	"fyne.io/fyne/v2/theme"
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
	CaCertPath := ""
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
	loginTab.Icon = theme.LoginIcon()
	// Add the login tab to the tabs container
	tabs.Append(loginTab)
	return inputBoxContainer
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
