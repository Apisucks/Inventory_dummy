package handlers

import (
	"Inventory_dummy/internal/models"
	"encoding/json"
	"net/http"
)

func DeleteHandler(rw http.ResponseWriter, r *http.Request) {
	nameOfObject := r.URL.Query().Get("name")
	if nameOfObject == "" {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode("Nothing provided for deletion")
		return
	}
	var temp []models.Object

	for _, v := range models.Objects {
		if v.Name == nameOfObject {
			continue
		}
		temp = append(temp, v)
	}
	models.Objects = temp

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode("Deletion successful")

}
