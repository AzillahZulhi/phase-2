package handler

import (
	"encoding/json"
	"net/http"
	"preview-week-2/config"
	"preview-week-2/entity"

	"github.com/julienschmidt/httprouter"
)

func GetByID(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var Branch entity.Branch

	err := config.DB.QueryRow("SELECT branch_id, name, location FROM branches WHERE branch_id = ?", id).
		Scan(&Branch.Branch_ID, &Branch.Name, &Branch.Location)
	if err != nil {
		http.Error(w, "Branch not found", http.StatusNotFound)
		return
	}

	SuccessMsg := entity.SuccessMessage{
		Message: "Show Data branch By ID",
		Status:  http.StatusOK,
		Data:    &Branch,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessMsg)
}

func GetAll(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	var Branches []entity.Branch

	rows, err := config.DB.Query("SELECT branch_id, name, location FROM branches")
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
		var Branch entity.Branch
		err := rows.Scan(&Branch.Branch_ID, &Branch.Name, &Branch.Location)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.ErrorMessage{
				Message: "Failed to scan result",
				Status:  http.StatusInternalServerError,
			})
			return
		}
		Branches = append(Branches, Branch)
	}

	var success_msg = entity.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Show data branches",
		Datas:   Branches,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success_msg)
}
