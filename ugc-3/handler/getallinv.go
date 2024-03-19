package handler

import (
	"encoding/json"
	"net/http"
	"ugc-3/config"
	"ugc-3/entity"

	"github.com/julienschmidt/httprouter"
)

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var inventories []entity.Inventory

	rows, err := config.DB.Query("SELECT id, name, stock, description, status FROM inventories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var inv entity.Inventory
		err := rows.Scan(&inv.ID, &inv.Name, &inv.Stock, &inv.Description, &inv.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		inventories = append(inventories, inv)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventories)
}
