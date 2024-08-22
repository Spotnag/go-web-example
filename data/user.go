package data

import (
	"github.com/google/uuid"
	"go-web-example/models"
	"golang.org/x/crypto/bcrypt"
)

func (d *Service) CreateUser(username, password string) (user *models.User, err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user = &models.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: string(hashedPassword),
	}
	_, err = d.DB.Exec("insert into user (id, username, password) values (?, ?, ?)", user.ID, user.Username, user.Password)
	return
}

func (d *Service) GetUser(username string) (user *models.User, err error) {
	row := d.DB.QueryRow("select id, username, password from user where username = ?", username)
	user = &models.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	return
}
