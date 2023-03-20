package aigenUi

import (
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func GenSettings(mapungubwe fyne.App) *container.TabItem {
	settingsTab := container.NewTabItem("Settings", widget.NewAccordion(

		widget.NewAccordionItem("Ai Voice Reply", widget.NewCheck("Audio Replies", func(OnandOff bool) {
			//TODO: Add a function to toggle the audio replies
			log.Printf("Audio Replies: %v", OnandOff)
			var soundIsOn = 1
			var soundIsOff = 0
			if OnandOff {
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

				// Insert the keylogger into the database
				_, err = db.Exec("INSERT INTO settings (audioOnly) VALUES (?)", soundIsOn)
				if err != nil {
					log.Printf("Error inserting into database: %v", err)
				}

			} else {
				//Store the keylogger in a file
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

				// Insert the keylogger into the database
				_, err = db.Exec("INSERT INTO settings (audioOnly) VALUES (?)", soundIsOff)
				if err != nil {
					log.Printf("Error inserting into database: %v", err)
				}

			}

		})),
		widget.NewAccordionItem("Add Watchlist", widget.NewLabel("Add a new stock to the watchlist")),
	),
	)
	return settingsTab
}
