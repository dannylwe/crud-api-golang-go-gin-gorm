package models

// Person is a datastruct used to hold person info
type Person struct {
	ID        int    `gorm:"AUTO_INCREMENT" json:"id"`
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
}
