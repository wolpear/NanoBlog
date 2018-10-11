package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

var webServerAddr = ":8000"

func init() {
	// Load ENVs
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// DB
	database = connectToDb()
	initDB()

	fmt.Printf("Starting NanoBlog\n")
}

func main() {
	defer database.Close()

	router := getRouter()
	s := &http.Server{
		Addr:           webServerAddr,
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
