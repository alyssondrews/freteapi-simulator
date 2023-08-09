package main

import (
	"log"
	"net/http"
)

func main() {

	database, err := db.Initialize()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	defer database.Close()

	http.HandleFunc("/quote", api.HandleQuote)
	http.HandleFunc("/metrics", api.HandleMetrics)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
