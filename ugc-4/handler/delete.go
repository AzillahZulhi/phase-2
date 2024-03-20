package handler

import (
	"net/http"
	"ugc-4/config"

	"github.com/julienschmidt/httprouter"
)

func DeleteReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	_, err := config.DB.Exec("DELETE FROM criminal_report WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
