package handlers

import (
	"Inventory_dummy/internal/models"
	"encoding/json"
	"net/http"
)

func AddHandler(rw http.ResponseWriter, r *http.Request) {
	var obj models.Object
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(rw, "failed to unmarshal the request", http.StatusBadRequest)
		return
	}
	models.Objects = append(models.Objects, obj)
	rw.WriteHeader(http.StatusOK)
}
