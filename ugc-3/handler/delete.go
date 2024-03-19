package handler

import (
	"net/http"
	"ugc-3/config"

	"github.com/julienschmidt/httprouter"
)

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	_, err := config.DB.Exec("DELETE FROM inventories WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
