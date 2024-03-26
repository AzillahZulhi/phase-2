package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"ugc-5/config"
	"ugc-5/entity"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newUser entity.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		})
		return
	}

	if newUser.Email == "" || !isValidEmail(newUser.Email) {
		errorMessage := entity.ErrorMessage{
			Message: "Invalid or empty email address",
			Status:  http.StatusBadRequest,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if newUser.Password == "" || len(newUser.Password) < 8 {
		errorMessage := entity.ErrorMessage{
			Message: "Password must be at least 8 characters long",
			Status:  http.StatusBadRequest,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if newUser.FullName == "" || len(newUser.FullName) < 6 || len(newUser.FullName) > 15 {
		errorMessage := entity.ErrorMessage{
			Message: "Full name must be between 6 and 15 characters long",
			Status:  http.StatusBadRequest,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if newUser.Age <= 17 {
		errorMessage := entity.ErrorMessage{
			Message: "Age must be at least 17",
			Status:  http.StatusBadRequest,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if newUser.Occupation == "" {
		errorMessage := entity.ErrorMessage{
			Message: "Occupation cannot be empty",
			Status:  http.StatusBadRequest,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if newUser.Role != "admin" && newUser.Role != "superadmin" {
		errorMessage := entity.ErrorMessage{
			Message: "Invalid role provided",
			Status:  http.StatusBadRequest,
		}
		http.Error(w, errorMessage.Message, errorMessage.Status)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	var count int
	err = config.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", newUser.Email).Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to check email uniqueness",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	if count > 0 {
		errorMessage := entity.ErrorMessage{
			Message: "Email already exists",
			Status:  http.StatusBadRequest,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to hash password",
			Status:  http.StatusBadRequest,
		})
		return
	}

	_, err = config.DB.Exec("INSERT INTO users (email, password, full_name, age, occupation, role) VALUES (?, ?, ?, ?, ?, ?)",
		newUser.Email, string(hashedPassword), newUser.FullName, newUser.Age, newUser.Occupation, newUser.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Failed to insert data into database",
			Status:  http.StatusBadRequest,
		})
		return
	}

	successMsg := entity.SuccessMessage{
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
