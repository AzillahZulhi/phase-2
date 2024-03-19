package handler

import (
	"UGC-2-3/config"
	"UGC-2-3/entity"
	"encoding/json"
	"net/http"
)

func GetHeroesData(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM Heroes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var heroes []entity.Hero
	for rows.Next() {
		var hero entity.Hero

		err = rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		heroes = append(heroes, hero)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}

func GetVillainsData(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM villain")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var villains []entity.Villain
	for rows.Next() {
		var villain entity.Villain

		err = rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.ImageURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		villains = append(villains, villain)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(villains)
}
