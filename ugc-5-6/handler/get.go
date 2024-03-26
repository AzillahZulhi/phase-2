package handler

import (
	"encoding/json"
	"net/http"
	"ugc-5/config"
	"ugc-5/entity"

	"github.com/julienschmidt/httprouter"
)

func GetByID(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var Recipe entity.Recipe

	err := config.DB.QueryRow("SELECT id, name, description FROM recipes WHERE id = ?", id).
		Scan(&Recipe.ID, &Recipe.Name, &Recipe.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Recipe not found",
			Status:  http.StatusBadRequest,
		})
		return
	}

	successMsg := entity.SuccessMessageR{
		Message: "Show Data Recipe By ID",
		Status:  http.StatusOK,
		Data:    &Recipe,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(successMsg)
}

func GetAll(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	var Recipes []entity.Recipe

	rows, err := config.DB.Query("SELECT id, name, description FROM recipes")
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
		var Recipe entity.Recipe
		err := rows.Scan(&Recipe.ID, &Recipe.Name, &Recipe.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.ErrorMessage{
				Message: "Failed to scan result",
				Status:  http.StatusInternalServerError,
			})
			return
		}
		Recipes = append(Recipes, Recipe)
	}

	var success_msg = entity.SuccessMessageR{
		Status:  http.StatusOK,
		Message: "Show Data recipes",
		Datas:   Recipes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success_msg)
}
