package auth

import (
	"sirclo/entities"
)

type Auth interface {
	Login(email string) (string, entities.User, error)
	GetEncryptPassword(email string) (string, error)
}
