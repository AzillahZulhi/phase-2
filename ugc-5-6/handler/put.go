package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"ugc-5/config"
	"ugc-5/entity"

	"github.com/julienschmidt/httprouter"
)

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	RecipeID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Invalid Recipe ID",
			Status:  http.StatusBadRequest,
		})
		return
	}

	var Recipe entity.Recipe
	err = json.NewDecoder(r.Body).Decode(&Recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Memperbarui data karyawan berdasarkan ID yang diberikan
	_, err = config.DB.Exec("UPDATE recipes SET name = ?, description = ?", Recipe.Name, Recipe.Description, RecipeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to update recipe data",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	successMsg := entity.SuccessMessageR{
		Message: "Successfully updated recipe with ID: " + id,
		Status:  http.StatusOK,
	}

	json.NewEncoder(w).Encode(successMsg)
}
