package models

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
