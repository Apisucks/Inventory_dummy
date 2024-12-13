package handlers

import (
	"Inventory_dummy/internal/models"
	"encoding/json"
	"net/http"
)

func ListHandler(rw http.ResponseWriter, r *http.Request) {

	if err := json.NewEncoder(rw).Encode(models.Objects); err != nil {
		http.Error(rw, "failed to marshal the response", http.StatusInternalServerError)
		return
	}

}
