package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type HealthCheckResponse struct {
	Database string `json:"database"`
}

func main() {
	r := gin.Default()

	// Define the health check route
	r.GET("/", healthCheck)

	// Start the server
	err := r.Run() // Default listens on :8080
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}

func healthCheck(c *gin.Context) {
	// Retrieve the database URL from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
	}

	// Try to connect to the database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"database": "Error"})
		return
	}
	defer db.Close()

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"database": "Error"})
		return
	}

	// Return a success response if the database is OK
	response := HealthCheckResponse{Database: "OK"}
	c.JSON(http.StatusOK, response)
}
