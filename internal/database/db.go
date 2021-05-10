package database

import "github.com/pcherednichenko/users/internal/models"

type DB interface {
	Get(id int, result *models.User) error
	Create(value *models.User) error
	Update(id int, value *models.User) error
	Delete(id int) error
	GetWithFilters(user models.User, users *[]models.User) error
}
