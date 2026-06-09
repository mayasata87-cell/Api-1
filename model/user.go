package model

import "gorm.io/gorm"

// Example model
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
