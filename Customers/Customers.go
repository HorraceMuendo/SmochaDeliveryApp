package customer

import "gorm.io/gorm"

type Details struct {
	gorm.Model

	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     uint   `json:"phone"`
	Location  string `json:"location"`
	Password  string `"json:"password"`
}
