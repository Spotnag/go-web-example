package data

import (
	"github.com/google/uuid"
	"go-web-example/models"
)

func (d *Service) CreateUserGroup(name string) (*models.UserGroup, error) {
	userGroup := &models.UserGroup{
		ID:   uuid.New(),
		Name: name,
	}

	_, err := d.DB.Exec("insert into usergroup (id, name) values (?, ?)", userGroup.ID, userGroup.Name)
	if err != nil {
		return nil, err
	}
	return userGroup, nil
}

func (d *Service) GetUserGroup(name string) (*models.UserGroup, error) {
	row := d.DB.QueryRow("select id, name from usergroup where name = ?", name)
	userGroup := &models.UserGroup{}

	err := row.Scan(&userGroup.ID, &userGroup.Name)
	if err != nil {
		return nil, err
	}
	return userGroup, nil
}
