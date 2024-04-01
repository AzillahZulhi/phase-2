package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"preview-week-3/config"
	"preview-week-3/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginCred entity.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&loginCred)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessageU{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		})
		return
	}

	var user entity.User
	row := config.DB.QueryRow("SELECT id, username, email, password, age, created_at, updated_at FROM users WHERE email = ?", loginCred.Email)
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			json.NewEncoder(w).Encode(entity.ErrorMessageU{
				Message: "Email not found",
				Status:  http.StatusBadRequest,
			})
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.ErrorMessageU{
				Message: "Internal Server error",
				Status:  http.StatusBadRequest,
			})
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginCred.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		json.NewEncoder(w).Encode(entity.ErrorMessageU{
			Message: "Incorrect password",
			Status:  http.StatusBadRequest,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
		"age":      user.Age,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte("mySecretKey")

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessageU{
			Message: "Failed to create token",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	successMsg := entity.SuccessMessageU{
		Message: "Login successful",
		Status:  http.StatusOK,
		Token:   tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(successMsg)

}
