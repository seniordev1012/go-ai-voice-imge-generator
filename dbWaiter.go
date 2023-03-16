package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
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
