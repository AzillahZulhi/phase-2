package handler

import (
	"encoding/json"
	"net/http"
	"ugc-4/config"
	"ugc-4/entity"

	"github.com/julienschmidt/httprouter"
)

func GetCriminalReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var Report entity.CriminalReport

	err := config.DB.QueryRow("SELECT id, hero_id, villain_id, description, date_of_incident, time_of_incident FORM criminal_report WHERE id = ?", id).Scan(&Report.ID, &Report.HeroID, &Report.VillainID, &Report.Description, &Report.Date, &Report.Time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Report)
}
