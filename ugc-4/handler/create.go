package handler

import (
	"encoding/json"
	"net/http"
	"ugc-4/config"
	"ugc-4/entity"

	"github.com/julienschmidt/httprouter"
)

func CreateCriminalReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var NewReport entity.CriminalReport

	err := json.NewDecoder(r.Body).Decode(&NewReport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("INSERT INTO criminal_report (hero_id, villain_id, description, date_of_incident, time_of_incident) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
