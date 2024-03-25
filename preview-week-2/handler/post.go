package handler

import (
	"encoding/json"
	"net/http"
	"preview-week-2/config"
	"preview-week-2/entity"

	"github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var NewBranch entity.Branch
	err := json.NewDecoder(r.Body).Decode(&NewBranch)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if NewBranch.Name == "" || NewBranch.Location == "" {
		http.Error(w, "Name and location cannot be empty", http.StatusBadRequest)
	}

	result, err := config.DB.Exec("INSERT INTO branches (name, location) VALUES (?, ?)", NewBranch.Name, NewBranch.Location)
	if err != nil {
		http.Error(w, "Failed to insert branch data", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
		return
	}

	NewBranch.Branch_ID = id
	successMsg := entity.SuccessMessage{
		Message: "Successfully created new branch",
		Status:  http.StatusCreated,
		Data:    &NewBranch,
	}
	jsonResponse, err := json.Marshal(successMsg)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
