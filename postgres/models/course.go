package models

type Course struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"artist"`
	Cost        int    `json:"cost"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
