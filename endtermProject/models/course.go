package models

type Course struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Cost        int    `json:"cost"`
}
