package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

/* ConnectDB establishes a connection to the database using environment variables */
func ConnectDB() {

	dbUser := DB_USER         // Get the database user from environment variables config.go
	dbPassword := DB_PASSWORD // Get the database password from environment variables config.go
	dbHost := DB_HOST         // Get the database host from environment variables config.go
	dbPort := DB_PORT         // Get the database port from environment variables config.go
	dbName := DB_NAME         // Get the database name from environment variables config.go

	// Create the DSN (Data Source Name) for MySQL connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the database
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	// Set connection pool settings
	DB.SetConnMaxLifetime(time.Minute * 5) // Set the maximum amount of time a connection may be reused (Time : 5 minutes)
	DB.SetConnMaxIdleTime(time.Minute * 2) // Set the maximum amount of time a connection may be idle (Time : 2 minutes)
	DB.SetMaxOpenConns(10)                 // Set the maximum number of open connections (Connection : 10)
	DB.SetMaxIdleConns(10)                 // Set the maximum number of idle connections (Connection : 10)

	err = DB.Ping() // Ping the database to verify the connection
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	log.Println("Successfully connected to the database")
}
