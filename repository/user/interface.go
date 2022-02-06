package user

import (
	"sirclo/entities"
)

type User interface {
	GetUsers() ([]entities.User, error)
	GetUserById(id int) (entities.User, error)
	CreateUser(user entities.User) error
	DeleteUser(id int) error
	EditUser(user entities.User, id int) error
}
