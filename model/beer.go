package model

type Beer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Strength    uint    `json:"strength"`
	AddedBy     string `json:"added_by"`
}
