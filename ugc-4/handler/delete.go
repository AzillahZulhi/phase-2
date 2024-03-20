package handler

import (
	"encoding/json"
	"net/http"
	"ugc-4/config"
	"ugc-4/entity"

	"github.com/julienschmidt/httprouter"
)

func DeleteReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var cr entity.CriminalReport
	err := json.NewDecoder(r.Body).Decode(&cr)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		})
		return
	}

	id := ps.ByName("id")
	statement, err := config.DB.Prepare("DELETE FROM criminal_report WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to prepare delete statement",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to execute delete statement",
			Status:  http.StatusBadRequest,
		})
		return
	}

	var successMsg = entity.SuccessMessage{
		Message: "Success: report deleted from id = " + id,
		Status:  http.StatusOK,
	}
	json.NewEncoder(w).Encode(successMsg)
}
