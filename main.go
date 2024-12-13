package main

import (
	"Inventory_dummy/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/inventory/list", handlers.ListHandler)
	server := http.Server{
		Addr:    ":9898",
		Handler: router,
	}

	log.Fatalln(server.ListenAndServe())
}
