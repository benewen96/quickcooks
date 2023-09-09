package models

type User struct {
	Entity
	Name            string
	Email           string
	Password        string `json:"-"`
	RoleAssignments []RoleAssignment
}
