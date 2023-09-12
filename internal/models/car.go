package models

type Car struct {
	Id         int    `json:"id"`
	Fabricator string `json:"fabricator"`
	Model      string `json:"model"`
	Year       int    `json:"year"`
}
