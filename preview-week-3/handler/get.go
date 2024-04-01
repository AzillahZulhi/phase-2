package handler

import (
	"encoding/json"
	"net/http"
	"preview-week-3/config"
	"preview-week-3/entity"

	"github.com/julienschmidt/httprouter"
)

func GetAllPhoto(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	var Photos []entity.Photo
	rows, err := config.DB.Query("SELECT title, caption, photo_url, user_id FROM photo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessageP{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var Photo entity.Photo
		err := rows.Scan(&Photo.Title, Photo.Caption, Photo.PhotoURL, Photo.User_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.ErrorMessageP{
				Message: "Failed to scan result",
				Status:  http.StatusInternalServerError,
			})
			return
		}
		Photos = append(Photos, Photo)
	}
	var success_msg = entity.SuccessMessageP{
		Status:  http.StatusOK,
		Message: "Show Data Photos",
		Datas:   Photos,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success_msg)
}

func GetAllComment(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	var Comments []entity.Comment
	rows, err := config.DB.Query("SELECT user_id, photo_id, message FROM comment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessageC{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var Comment entity.Comment
		err := rows.Scan(&Comment.User_id, &Comment.Photo_id, &Comment.Message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.ErrorMessageC{
				Message: "Failed to scan result",
				Status:  http.StatusInternalServerError,
			})
			return
		}
		Comments = append(Comments, Comment)
	}
	var success_msg = entity.SuccessMessageC{
		Status:  http.StatusOK,
		Message: "Show Data Comments",
		Datas:   Comments,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success_msg)
}

func GetAllSM(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	var SMs []entity.SocialMedia
	rows, err := config.DB.Query("SELECT name, social_media_url, user_id FROM social_media")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessageS{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var SM entity.SocialMedia
		err := rows.Scan(&SM.Name, &SM.Social_media_url, &SM.User_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.ErrorMessageS{
				Message: "Failed to scan result",
				Status:  http.StatusInternalServerError,
			})
			return
		}
		SMs = append(SMs, SM)
	}
	var success_msg = entity.SuccessMessageS{
		Status:  http.StatusOK,
		Message: "Show Data Social medias",
		Datas:   SMs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success_msg)
}
