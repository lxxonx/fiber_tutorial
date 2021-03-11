package models

type User struct {
	ID 			uint	`json:"id" gorm:"unique"`
	Name 		string	`json:"name"`
	Email 		string 	`json:"email" gorm:"unique"`
	Password 	[]byte	`json:"-"` // don't return the value
}