package handler

import (
	"encoding/json"
	"net/http"
	"ugc-3/config"
	"ugc-3/entity"

	"github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var inv entity.Inventory

	err := json.NewDecoder(r.Body).Decode(&inv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("INSERT INTO inventories (name, stock, description, status) VALUES (?, ?, ?, ?)", inv.Name, inv.Stock, inv.Description, inv.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
