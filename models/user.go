package models

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password"`
	RoleID       uuid.UUID `db:"role_id"`
	Role         *Role
}

type Role struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Permissions []string  `db:"permissions"`
}

const (
	PermissionManageUsers        = "ManageUsers"
	PermissionManageCourses      = "ManageCourses"
	PermissionAssignCourses      = "AssignCourses"
	PermissionResetCredentials   = "ResetCredentials"
	PermissionBulkUploadUsers    = "BulkUploadUsers"
	PermissionManageGroupUsers   = "ManageGroupUsers"
	PermissionAssignGroupCourses = "AssignGroupCourses"
	PermissionViewCourses        = "ViewCourses"
)
