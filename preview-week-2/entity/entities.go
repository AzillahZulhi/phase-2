package entity

type Branch struct {
	Branch_ID int64  `json:"branch_id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
}

type ErrorMessage struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
}

type SuccessMessage struct {
	Message string   `json:"message"`
	Status  int64    `json:"status"`
	Datas   []Branch `json:"datas,omitempty"`
	Data    *Branch  `json:"data,omitempty"`
}
