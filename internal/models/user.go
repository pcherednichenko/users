package models

// User model with first name, last name, nickname, password, email and country
type User struct {
	ID        int `gorm:"primaryKey;->:false;<-:create"` // read only
	FirstName string
	LastName  string
	Nickname  string
	Password  string `gorm:"->:false"` // create and update only (disabled read from db)
	Email     string
	Country   string
}
