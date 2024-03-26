package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
	"ugc-5/config"
	"ugc-5/entity"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("12345")

func LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginCred entity.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&loginCred)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		})
		return
	}

	query := "SELECT id, email, password, full_name, age, occupation, role FROM users WHERE email = ?"
	row := config.DB.QueryRow(query, loginCred.Email)

	var user entity.User
	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.FullName, &user.Age, &user.Occupation, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			json.NewEncoder(w).Encode(entity.ErrorMessage{
				Message: "Email not found",
				Status:  http.StatusBadRequest,
			})
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.ErrorMessage{
				Message: "Internal Server error",
				Status:  http.StatusBadRequest,
			})
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginCred.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Incorrect password",
			Status:  http.StatusBadRequest,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorMessage{
			Message: "Error signing token",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	successMsg := entity.SuccessMessage{
		Message: "Login successful",
		Status:  http.StatusOK,
		Token:   tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(successMsg)
}
