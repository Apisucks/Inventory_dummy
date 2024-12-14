package main

import (
	"Inventory_dummy/internal/handlers"
	"Inventory_dummy/internal/models"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func init() {
	models.LoadConfig()
}

func main() {

	router := mux.NewRouter()

	version := fmt.Sprintf("%v", models.ConfigMap["version"])
	if strings.TrimSpace(version) == "" {
		panic(fmt.Errorf("service version not found"))
	}
	router.HandleFunc("/inventory/"+version+"/add", handlers.AddHandler).Methods(http.MethodPost)
	router.HandleFunc("/inventory/"+version+"/list", handlers.ListHandler).Methods(http.MethodGet)
	router.HandleFunc("/inventory/"+version+"/delete", handlers.DeleteHandler).Methods(http.MethodGet)
	server := http.Server{
		Addr:    ":9898",
		Handler: router,
	}

	log.Fatalln(server.ListenAndServe())
}
