package handler

import (
	"encoding/json"
	"net/http"
	"ugc-5/config"
	"ugc-5/entity"

	"github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var NewRecipe entity.Recipe
	err := json.NewDecoder(r.Body).Decode(&NewRecipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if NewRecipe.Name == "" || NewRecipe.Description == "" {
		errorMessage := entity.ErrorMessage{
			Message: "Name or description cannot be empty",
			Status:  http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	result, err := config.DB.Exec("INSERT INTO recipes (name, description) VALUES (?, ?)",
		NewRecipe.Name, NewRecipe.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to insert data into database",
			Status:  http.StatusBadRequest,
		})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to retrieve last insert ID",
			Status:  http.StatusBadRequest,
		})
		return
	}

	NewRecipe.ID = int(id)
	successMsg := entity.SuccessMessageR{
		Message: "Successfully created new recipe",
		Status:  http.StatusCreated,
		Data:    &NewRecipe,
	}

	jsonResponse, err := json.Marshal(successMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to marshal JSON response",
			Status:  http.StatusBadRequest,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
