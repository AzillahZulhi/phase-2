package handler

import (
	"encoding/json"
	"net/http"
	"ugc-4/config"
	"ugc-4/entity"

	"github.com/julienschmidt/httprouter"
)

func GetALLCriminalReport(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var Reports []entity.CriminalReport

	rows, err := config.DB.Query("SELECT * FROM criminal_report")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var Report entity.CriminalReport
		err := rows.Scan(&Report.ID, &Report.HeroID, &Report.VillainID, &Report.Description, &Report.Date, &Report.Time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Reports = append(Reports, Report)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Reports)
	}
}
