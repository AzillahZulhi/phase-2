package handler

import (
	"encoding/json"
	"net/http"
	"preview-week-3/config"
	"preview-week-3/entity"
	"regexp"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newUser entity.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageU{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
			Data:    &newUser,
		})
		return
	}

	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" {
		errorMessage := entity.ErrorMessageU{
			Message: "Username, email, and password are required",
			Status:  http.StatusBadRequest,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	EmailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !EmailRegex.MatchString(newUser.Email) {
		errorMessage := entity.ErrorMessageU{
			Message: "Invalid email format",
			Status:  http.StatusBadRequest,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if len(newUser.Password) < 8 {
		errorMessage := entity.ErrorMessageU{
			Message: "Password must be at least 8 characters long",
			Status:  http.StatusBadRequest,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageU{
			Message: "Failed to hash password",
			Status:  http.StatusBadRequest,
		})
		return
	}

	_, err = config.DB.Exec("INSERT INTO users (username, email, password, age) VALUES (?, ?, ?, ?)",
		newUser.Username, newUser.Email, string(hashedPassword), newUser.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageU{
			Message: "Failed to insert data into database",
			Status:  http.StatusBadRequest,
		})
		return
	}

	successMsg := entity.SuccessMessageU{
		Message: "Successfully Registered",
		Status:  http.StatusCreated,
		Data:    &newUser,
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
