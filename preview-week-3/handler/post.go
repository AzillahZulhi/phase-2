package handler

import (
	"encoding/json"
	"net/http"
	"preview-week-3/config"
	"preview-week-3/entity"

	"github.com/julienschmidt/httprouter"
)

func CreatePhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var NewPhoto entity.Photo
	err := json.NewDecoder(r.Body).Decode(&NewPhoto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageP{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
			Data:    &NewPhoto,
		})
		return
	}

	if NewPhoto.Title == "" || NewPhoto.Caption == "" || NewPhoto.PhotoURL == "" {
		errorMessage := entity.ErrorMessageP{
			Message: "Title, Caption, and PhotoURL cannot be empty",
			Status:  http.StatusBadRequest,
			Data:    &NewPhoto,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if NewPhoto.User_id <= 0 {
		errorMessage := entity.ErrorMessageP{
			Message: "Invalid user ID",
			Status:  http.StatusBadRequest,
			Data:    &NewPhoto,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	tx, err := config.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	result, err := tx.Exec("INSERT INTO photo (title, caption, photo_url, user_id) VALUES (?, ?, ?, ?)",
		NewPhoto.Title, NewPhoto.Caption, NewPhoto.PhotoURL, NewPhoto.User_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageP{
			Message: "Failed to insert data into database",
			Status:  http.StatusBadRequest,
		})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageP{
			Message: "Failed to retrieve last insert ID",
			Status:  http.StatusBadRequest,
		})
		return
	}

	NewPhoto.ID = int(id)
	successMsg := entity.SuccessMessageP{
		Message: "Successfully created new photo",
		Status:  http.StatusCreated,
		Data:    &NewPhoto,
	}

	jsonResponse, err := json.Marshal(successMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageP{
			Message: "Failed to marshal JSON response",
			Status:  http.StatusBadRequest,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateComment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var NewComment entity.Comment
	err := json.NewDecoder(r.Body).Decode(&NewComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageC{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
			Data:    &NewComment,
		})
		return
	}

	if NewComment.User_id <= 0 || NewComment.Photo_id <= 0 {
		errorMessage := entity.ErrorMessageC{
			Message: "Invalid user ID or photo ID",
			Status:  http.StatusBadRequest,
			Data:    &NewComment,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if NewComment.Message == "" {
		errorMessage := entity.ErrorMessageC{
			Message: "Message cannot be empty",
			Status:  http.StatusBadRequest,
			Data:    &NewComment,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	tx, err := config.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	result, err := tx.Exec("INSERT INTO comment (user_id, photo_id, message) VALUES (?, ?, ?)",
		NewComment.User_id, NewComment.Photo_id, NewComment.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageP{
			Message: "Failed to insert data into database",
			Status:  http.StatusBadRequest,
		})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageP{
			Message: "Failed to retrieve last insert ID",
			Status:  http.StatusBadRequest,
		})
		return
	}

	NewComment.ID = int(id)
	successMsg := entity.SuccessMessageC{
		Message: "Successfully created new comment",
		Status:  http.StatusCreated,
		Data:    &NewComment,
	}

	jsonResponse, err := json.Marshal(successMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageC{
			Message: "Failed to marshal JSON response",
			Status:  http.StatusBadRequest,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateSocialMedia(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var NewSM entity.SocialMedia
	err := json.NewDecoder(r.Body).Decode(&NewSM)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageS{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
			Data:    &NewSM,
		})
		return
	}

	if NewSM.Name == "" || NewSM.Social_media_url == "" {
		errorMessage := entity.ErrorMessageS{
			Message: "Name ad url canntot be empty",
			Status:  http.StatusBadRequest,
			Data:    &NewSM,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if NewSM.User_id <= 0 {
		errorMessage := entity.ErrorMessageS{
			Message: "Invalid user ID",
			Status:  http.StatusBadRequest,
			Data:    &NewSM,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	tx, err := config.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	result, err := tx.Exec("INSERT INTO social_media (name, social_media_url, user_id) VALUES (?, ?, ?)",
		NewSM.Name, NewSM.Social_media_url, NewSM.User_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageS{
			Message: "Failed to insert data into database",
			Status:  http.StatusBadRequest,
		})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageP{
			Message: "Failed to retrieve last insert ID",
			Status:  http.StatusBadRequest,
		})
		return
	}

	NewSM.ID = int(id)
	successMsg := entity.SuccessMessageS{
		Message: "Successfully created new photo",
		Status:  http.StatusCreated,
		Data:    &NewSM,
	}

	jsonResponse, err := json.Marshal(successMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageP{
			Message: "Failed to marshal JSON response",
			Status:  http.StatusBadRequest,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
