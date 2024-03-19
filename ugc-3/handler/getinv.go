package handler

import (
	"encoding/json"
	"net/http"
	"ugc-3/config"
	"ugc-3/entity"

	"github.com/julienschmidt/httprouter"
)

func GetInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var inv entity.Inventory

	err := config.DB.QueryRow("SELECT id, name, stock, description, status FROM inventories WHERE id = ?", id).Scan(&inv.ID, &inv.Name, &inv.Stock, &inv.Description, &inv.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inv)
}
