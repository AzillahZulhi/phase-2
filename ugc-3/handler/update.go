package handler

import (
	"encoding/json"
	"net/http"
	"ugc-3/config"
	"ugc-3/entity"

	"github.com/julienschmidt/httprouter"
)

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var inv entity.Inventory

	err := json.NewDecoder(r.Body).Decode(&inv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("UPDATE inventories SET name = ?, stock = ?, description = ?, status = ? WHERE id = ?", inv.Name, inv.Stock, inv.Description, inv.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
