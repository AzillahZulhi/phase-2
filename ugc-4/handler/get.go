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

	rows, err := config.DB.Query("SELECT id, hero_id, villain_id, description, date_of_incident, time_of_incident FROM criminal_report WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&Report.ID, &Report.HeroID, &Report.VillainID, &Report.Description, &Report.Date, &Report.Time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.ErrorMessage{
				Message: "Failed to scan result",
				Status:  http.StatusInternalServerError,
			})
			return
		}
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Error reading rows",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	successMsg := entity.SuccessMessage{
		Message: "Show Data Criminal Report By ID",
		Status:  http.StatusOK,
		Data:    &Report,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(successMsg)
}
