package handler

import (
	"encoding/json"
	"net/http"
	"preview-week-2/config"
	"preview-week-2/entity"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	BranchID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Invalid branch ID",
			Status:  http.StatusBadRequest,
		})
		return
	}

	var Branch entity.Branch
	err = json.NewDecoder(r.Body).Decode(&Branch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		})
		return
	}

	_, err = config.DB.Exec("UPDATE branches SET name = ?, location = ? WHERE branch_id = ?", Branch.Name, Branch.Location, BranchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to update branch data",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	successMsg := entity.SuccessMessage{
		Message: "Successfully updated branch with ID: " + id,
		Status:  http.StatusOK,
	}

	json.NewEncoder(w).Encode(successMsg)
}
