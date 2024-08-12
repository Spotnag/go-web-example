package service

import (
	"database/sql"
	"github.com/google/uuid"
	"go-web-example/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (u *UserService) CreateUser(username, password string) (user *models.User, err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user = &models.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: string(hashedPassword),
	}

	_, err = u.DB.Exec("insert into user (id, username, password) values (?, ?, ?)", user.ID, user.Username, user.Password)
	return
}

func (u *UserService) GetUser(username string) (user *models.User, err error) {
	row := u.DB.QueryRow("select id, username, password from user where username = ?", username)
	user = &models.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	return
}

func (u *UserService) LoginUser(username, password string) (*models.User, error) {
	targetUser, err := u.GetUser(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(targetUser.Password), []byte(password)); err != nil {
		return nil, err
	}
	return targetUser, nil
}
