package models

type User struct {
	ID        string `json:"id" xml:"id" gorm:"primaryKey"`
	Email     string `json:"email" xml:"email"`
	Password  string `json:"password" xml:"password"`
	SecretKey string `json:"secret_key" xml:"secret_key"`
	AuthID	  string `json:"auth_id" xml:"auth_id"`
}
