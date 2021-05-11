package handler

import "github.com/pcherednichenko/users/internal/models"

// MockDB for test
type MockDB struct {
	OnGet            func(id int, result *models.User) error
	OnCreate         func(value *models.User) error
	OnUpdate         func(id int, value *models.User) error
	OnDelete         func(id int) error
	OnGetWithFilters func(user models.User, users *[]models.User) error
}

func (m *MockDB) Get(id int, result *models.User) error {
	return m.OnGet(id, result)
}
func (m *MockDB) Create(value *models.User) error {
	return m.OnCreate(value)
}
func (m *MockDB) Update(id int, value *models.User) error {
	return m.OnUpdate(id, value)
}
func (m *MockDB) Delete(id int) error {
	return m.OnDelete(id)
}
func (m *MockDB) GetWithFilters(user models.User, users *[]models.User) error {
	return m.OnGetWithFilters(user, users)
}
