package entities

import "golang.org/x/crypto/bcrypt"

// dipisah folder pisah file juga
type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	ImageUrl string `json:"image_url" form:"image_url"`
}

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ComparePassword(hashedPassword string, loginPassword string) error {
	err_token := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginPassword))
	return err_token
}
