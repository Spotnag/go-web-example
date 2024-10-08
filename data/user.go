package data

import (
	"github.com/google/uuid"
	"go-web-example/models"
	"golang.org/x/crypto/bcrypt"
)

func (d *Service) CreateUser(email, password, role string) (user *models.User, err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	r, err := d.GetRole(role)
	if err != nil {
		return nil, err
	}

	user = &models.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: string(hashedPassword),
		RoleID:       r.ID,
	}
	_, err = d.DB.Exec("insert into user (id, email, password) values (?, ?, ?)", user.ID, user.Email, user.PasswordHash, user.RoleID)
	return
}

func (d *Service) GetUser(email string) (user *models.User, err error) {
	row := d.DB.QueryRow("select id, email, password from user where email = ?", email)
	user = &models.User{}
	err = row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	return
}
