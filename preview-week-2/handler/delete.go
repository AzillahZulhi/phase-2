package handler

import (
	"encoding/json"
	"net/http"
	"preview-week-2/config"
	"preview-week-2/entity"

	"github.com/julienschmidt/httprouter"
)

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	statement, err := config.DB.Prepare("DELETE FROM branches WHERE branch_id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	successMsg := entity.SuccessMessage{
		Message: "Success: Branch data deleted, id = " + id,
		Status:  http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(successMsg)
}
