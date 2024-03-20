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
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		})
		return
	}

	result, err := config.DB.Exec("INSERT INTO criminal_report (hero_id, villain_id, description, date_of_incident, time_of_incident) VALUES (?, ?, ?, ?, ?)", NewReport.HeroID, NewReport.VillainID, NewReport.Description, NewReport.Date, NewReport.Time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to create report",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to retrieve last insert ID",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	NewReport.ID = id
	successMsg := entity.SuccessMessage{
		Message: "Successfully created new report",
		Status:  http.StatusCreated,
		Data:    &NewReport,
	}

	jsonResponse, err := json.Marshal(successMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
