package database

import "github.com/pcherednichenko/users/internal/models"

// DB is a main interface with all needed functions for current logic
type DB interface {
	Get(id int, result *models.User) error
	Create(value *models.User) error
	Update(id int, value *models.User) error
	Delete(id int) error
	GetWithFilters(user models.User, users *[]models.User) error
}
