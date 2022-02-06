package user

import (
	"database/sql"
	"fmt"
	"sirclo/entities"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUsers() ([]entities.User, error) {
	var users []entities.User
	result, err := ur.db.Query("select id, name, email, image_url from users where deleted_at is null")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var user entities.User
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.ImageUrl)
		if err != nil {
			return nil, fmt.Errorf("user not found")
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetUserById(id int) (entities.User, error) {
	var user entities.User
	result, err := ur.db.Query("select id, name, email, image_url from users where deleted_at is null AND id = ?", id)
	if err != nil {
		return user, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.ImageUrl)
		if err != nil {
			return user, fmt.Errorf("user not found")
		}
		return user, nil
	}
	return user, fmt.Errorf("user not found")
}

func (ur *UserRepository) DeleteUser(id int) error {
	result, err := ur.db.Exec("UPDATE users SET deleted_at = now() where id = ? AND deleted_at IS null", id)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (ur *UserRepository) EditUser(user entities.User, id int) error {
	result, err := ur.db.Exec("UPDATE users SET name= ?, email= ?, password= ?, image_url = ?, updated_at = now() WHERE id = ? AND deleted_at IS null", user.Name, user.Email, user.Password, user.ImageUrl, id)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (ur *UserRepository) CreateUser(user entities.User) error {
	result, err := ur.db.Exec("INSERT INTO users(name, email, password) VALUES(?,?,?)", user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not created")
	}
	return nil
}
