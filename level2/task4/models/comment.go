package models

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
