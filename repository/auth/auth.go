package auth

import (
	"database/sql"
	"fmt"
	"sirclo/delivery/middlewares"
	"sirclo/entities"
)

type AuthRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) Login(email string) (string, entities.User, error) {
	var user entities.User
	// input email = asd, password = 123
	result, err := a.db.Query("select id, name, email from users where email = ?", email)
	if err != nil {
		fmt.Println(err)
		return "", user, err
	}
	for result.Next() {
		err_scan := result.Scan(&user.Id, &user.Name, &user.Email)
		if err_scan != nil {
			return "", user, err_scan
		}
	}
	if user.Email == email {
		token, err_token := middlewares.CreateToken(user.Id)
		if err_token != nil {
			return "", user, err
		}
		return token, user, nil
	}
	// tidak error tapi usernya tidak ada
	return "", user, fmt.Errorf("user not found")
}

func (a *AuthRepository) GetEncryptPassword(email string) (string, error) {
	var password string
	// input email = asd, password = 123
	result, err := a.db.Query("select password from users where email = ? AND deleted_at IS null", email)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for result.Next() {
		err_scan := result.Scan(&password)
		if err_scan != nil {
			return "", err_scan
		}
		return password, nil
	}
	return "", fmt.Errorf("User Not Found")
}
