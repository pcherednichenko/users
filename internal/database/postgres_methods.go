package database

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/pcherednichenko/users/internal/models"
)

// TODO: we are using standard functions from gorm
// but in real life better to test this with sql-mock

// Get user by id
func (p *PostgresDB) Get(id int, value *models.User) error {
	err := p.db.Model(&models.User{}).Where("ID = ?", id).First(value).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf(notFoundErr, id)
	}
	return err
}

// Create user in database
func (p *PostgresDB) Create(value *models.User) error {
	return p.db.Create(value).Error
}

// Update user by ad and new user data
func (p *PostgresDB) Update(id int, value *models.User) error {
	updResult := p.db.Model(&models.User{}).Where("ID = ?", id).Update(value)
	if updResult.RowsAffected == 0 {
		// TODO: such cases are not internal error, it is bad request
		return fmt.Errorf(notFoundErr, id)
	}
	return updResult.Error
}

// Delete user by id
func (p *PostgresDB) Delete(id int) error {
	updResult := p.db.Model(&models.User{}).Delete(&models.User{}, "ID = ?", id)
	if updResult.RowsAffected == 0 {
		return fmt.Errorf(notFoundErr, id)
	}
	return updResult.Error
}

// GetWithFilters get users with some filters
func (p *PostgresDB) GetWithFilters(user models.User, users *[]models.User) error {
	return p.db.Where(&user).Find(users).Error
}
