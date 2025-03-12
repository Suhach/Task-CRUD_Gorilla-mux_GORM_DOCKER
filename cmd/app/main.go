package main

import (
	"log"
	"net/http"
	"test_ByMyself/internal/database"
	"test_ByMyself/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	r := mux.NewRouter()
	handlers.RegisterTaskRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
