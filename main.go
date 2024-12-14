package main

import (
	"Inventory_dummy/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/inventory/add", handlers.AddHandler).Methods(http.MethodPost)
	router.HandleFunc("/inventory/list", handlers.ListHandler).Methods(http.MethodGet)
	router.HandleFunc("/inventory/delete", handlers.DeleteHandler).Methods(http.MethodGet)
	server := http.Server{
		Addr:    ":9898",
		Handler: router,
	}

	log.Fatalln(server.ListenAndServe())
}
