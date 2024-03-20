package entity

type Hero struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	ImageURL string `json:"image_url"`
}

type Villain struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	ImageURL string `json:"image_url"`
}

type CriminalReport struct {
	ID          int    `json:"id"`
	HeroID      int    `json:"hero_id"`
	VillainID   int    `json:"villain_id"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Time        string `json:"time"`
}
