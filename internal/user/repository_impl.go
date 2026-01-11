package user

import (
	"fmt"

	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

var (
	users = []User  {
		{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "12345"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Password: "12345"},
	}
)

func NewRepositoryImpl(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) FindAll() ([]User, error) {

	//horario := r.db.Raw("SELECT NOW()")

	return users, nil
}

func (r *RepositoryImpl) FindByID(id int) (*User, error) {
	
	for i := range users {
		if users[i].ID == id {
			return &users[i], nil
		}
	} 
	
	return nil, fmt.Errorf("There is no user with that ID")
}