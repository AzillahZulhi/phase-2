package handler

import (
	"encoding/json"
	"net/http"
	"ugc-5/config"
	"ugc-5/entity"

	"github.com/julienschmidt/httprouter"
)

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var Recipe entity.Recipe
	err := json.NewDecoder(r.Body).Decode(&Recipe)
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
	statement, err := config.DB.Prepare("DELETE FROM recipes WHERE id = ?")
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

	var successMsg = entity.SuccessMessageR{
		Message: "Success: Employee deleted, id = " + id,
		Status:  http.StatusOK,
	}
	json.NewEncoder(w).Encode(successMsg)
}
