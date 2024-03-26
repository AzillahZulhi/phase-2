package entity

type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FullName   string `json:"full_name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
	Role       string `json:"role"`
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Recipe struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ErrorMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type SuccessMessage struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Datas   []User `json:"datas,omitempty"`
	Data    *User  `json:"data,omitempty"`
	Token   string `json:"token,omitempty"`
}

type SuccessMessageR struct {
	Message string   `json:"message"`
	Status  int64    `json:"status"`
	Datas   []Recipe `json:"datas,omitempty"`
	Data    *Recipe  `json:"data,omitempty"`
}
