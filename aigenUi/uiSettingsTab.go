package aigenUi

import (
	"aigen/aigenRest"
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
)

const (
	setEnvironment = "setenv"
	saveToDB       = "savedb"
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

		})), addToWatchList(), addOpenAiKeys(), azureSpeechKeys()))
	settingsTab.Icon = theme.SettingsIcon()
	return settingsTab
}

func azureSpeechKeys() *widget.AccordionItem {
	return uiFormTemplate("Azure Speech Keys",
		setEnvironment,
		"Azure Speech",
		"Enter Your Azure Speech Keys",
		"SPEECH_KEY")
}

func addToWatchList() *widget.AccordionItem {
	return widget.NewAccordionItem("Add Watchlist", widget.NewLabel("Add a new stock to the watchlist"))
}

func addOpenAiKeys() *widget.AccordionItem {

	return uiFormTemplate("OpenAI Keys",
		setEnvironment,
		"OpenAI Keys",
		"Enter Your OpenAI Keys",
		"OPENAI")

}

func uiFormTemplate(buttonText string, trigger string, formAction string, placeHolder string, valueKey string) *widget.AccordionItem {

	var formValue string
	keysEntry := widget.NewEntry()
	apiTokens := formButton(buttonText, trigger, formValue, keysEntry, valueKey)

	form := formFields(placeHolder, keysEntry, apiTokens)
	keysEntry.OnChanged = func(s string) {
		log.Printf(s)
	}
	keysEntry.FocusGained()

	return widget.NewAccordionItem(formAction, form)
}

func formButton(buttonText string, trigger string, formValue string, keysEntry *widget.Entry, valueKey string) *widget.Button {
	apiTokens := widget.NewButton(buttonText, func() {
		formValue = keysEntry.Text
		//Insert Actions to different function : setenv or savedb
		switch trigger {
		case setEnvironment:
			setValueToEnv(formValue, valueKey)
		case saveToDB:
			saveValueToDB(formValue, valueKey)
		}

	})
	return apiTokens
}

func formFields(placeHolder string, keysEntry *widget.Entry, apiTokens *widget.Button) *widget.Form {
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: placeHolder, Widget: keysEntry},
		},
		OnSubmit: apiTokens.OnTapped,
		OnCancel: func() {
			log.Println("Cancelled")
		},
	}
	return form
}

func saveValueToDB(value string, key string) {

}

func setValueToEnv(value string, key string) {
	setValue := os.Setenv(key, value)
	if setValue != nil {
		log.Printf("Error setting environment variable: %v", setValue)
	}
	aigenRest.SendNotificationNow(key + "Saved")
	log.Printf("Environment setup: %v", key)
}
