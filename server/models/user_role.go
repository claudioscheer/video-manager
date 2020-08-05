package models

// UserRole database table.
type UserRole struct {
	User UserModel
	Role RoleModel
}
