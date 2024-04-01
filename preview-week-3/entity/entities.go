package entity

import "time"

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Age        int       `json:"age"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Photo struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	PhotoURL   string    `json:"photo_url"`
	User_id    int       `json:"user_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Comment struct {
	ID         int       `json:"id"`
	User_id    int       `json:"user_id"`
	Photo_id   int       `json:"photo_id"`
	Message    string    `json:"message"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type SocialMedia struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Social_media_url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	Created_at       time.Time `json:"created_at"`
	Updated_at       time.Time `json:"updated_at"`
}

type ErrorMessageU struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Datas   []User `json:"datas,omitempty"`
	Data    *User  `json:"data,omitempty"`
}

type ErrorMessageP struct {
	Message string  `json:"message"`
	Status  int     `json:"status"`
	Datas   []Photo `json:"datas,omitempty"`
	Data    *Photo  `json:"data,omitempty"`
}

type ErrorMessageS struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Datas   []SocialMedia `json:"datas,omitempty"`
	Data    *SocialMedia  `json:"data,omitempty"`
}

type ErrorMessageC struct {
	Message string    `json:"message"`
	Status  int       `json:"status"`
	Datas   []Comment `json:"datas,omitempty"`
	Data    *Comment  `json:"data,omitempty"`
}

type SuccessMessageU struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Datas   []User `json:"datas,omitempty"`
	Data    *User  `json:"data,omitempty"`
	Token   string `json:"token,omitempty"`
}

type SuccessMessageP struct {
	Message string  `json:"message"`
	Status  int64   `json:"status"`
	Datas   []Photo `json:"datas,omitempty"`
	Data    *Photo  `json:"data,omitempty"`
}

type SuccessMessageC struct {
	Message string    `json:"message"`
	Status  int64     `json:"status"`
	Datas   []Comment `json:"datas,omitempty"`
	Data    *Comment  `json:"data,omitempty"`
}

type SuccessMessageS struct {
	Message string        `json:"message"`
	Status  int64         `json:"status"`
	Datas   []SocialMedia `json:"datas,omitempty"`
	Data    *SocialMedia  `json:"data,omitempty"`
}
