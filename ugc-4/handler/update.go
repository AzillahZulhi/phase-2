package handler

import (
	"encoding/json"
	"net/http"
	"ugc-4/config"
	"ugc-4/entity"

	"github.com/julienschmidt/httprouter"
)

func UpdateCriminalReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var Report entity.CriminalReport

	err := json.NewDecoder(r.Body).Decode(&Report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		})
		return
	}

	_, err = config.DB.Exec("UPDATE criminal_report SET hero_id = ?, villain_id = ?, description = ?, date_of_incident = ?, time_of_incident = ? WHERE id = ?", Report.HeroID, Report.VillainID, Report.Description, Report.Date, Report.Time, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	successMsg := entity.SuccessMessage{
		Message: "success update report from id = " + id,
		Status:  http.StatusOK,
	}

	json.NewEncoder(w).Encode(successMsg)
}
