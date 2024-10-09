package data

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-web-example/models"
)

func (d *Service) CreateRole(name string, permissions []string) (*models.Role, error) {
	jsonPermissions, err := json.Marshal(permissions)
	if err != nil {
		return nil, err
	}

	role := &models.Role{
		ID:          uuid.New(),
		Name:        name,
		Permissions: permissions,
	}

	_, err = d.DB.Exec("insert into role (id, name, permissions) values (?, ?, ?)", role.ID, role.Name, string(jsonPermissions))
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (d *Service) GetRole(name string) (*models.Role, error) {
	row := d.DB.QueryRow("select id, name, permissions from role where name = ?", name)
	role := &models.Role{}
	var permissionsJSON string

	err := row.Scan(&role.ID, &role.Name, &permissionsJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(permissionsJSON), &role.Permissions)
	if err != nil {
		return nil, err
	}
	return role, nil
}
