package main

import (
	"log"
	"net/http"

	api "lesson/handlers"
	"lesson/storage"
)

func main() {
	db, err := storage.InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()

	r := api.SetupRouter(db)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
