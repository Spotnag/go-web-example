package data

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-web-example/models"
	"golang.org/x/crypto/bcrypt"
)

func (d *Service) CreateUser(email, password, role, usergroup string) (*models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	r, err := d.GetRole(role) // validate role exists before creating user
	if err != nil {
		return nil, err
	}

	u, err := d.GetUserGroup(usergroup) // validate usergroup exists before creating user
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: string(hashedPassword),
		RoleID:       r.ID,
		UserGroupID:  u.ID,
	}
	_, err = d.DB.Exec("insert into user (id, email, password, role_id, usergroup_id) values (?, ?, ?, ?, ?)", user.ID, user.Email, user.PasswordHash, user.RoleID, user.UserGroupID)
	return user, err
}

//func (d *Service) GetUser(email string) (user *models.User, err error) {
//	row := d.DB.QueryRow("select id, email, password, role_id, usergroup_id from user where email = ?", email)
//	user = &models.User{}
//	err = row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.RoleID)
//	return
//}

func (d *Service) GetUser(email string) (*models.User, error) {
	query := `
	   SELECT u.id, u.email, u.password, u.role_id, u.usergroup_id,
	          r.id, r.name, r.permissions
	   FROM user u
	   JOIN role r ON u.role_id = r.id
	   WHERE u.email = ?
	`

	row := d.DB.QueryRow(query, email)

	user := &models.User{}
	role := &models.Role{}
	var permissionsJSON string

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.RoleID, &user.UserGroupID, &role.ID, &role.Name, &permissionsJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(permissionsJSON), &role.Permissions)
	if err != nil {
		return nil, err
	}

	user.Role = role

	return user, nil
}
