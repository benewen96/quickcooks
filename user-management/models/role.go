package models

import "gorm.io/gorm"

type RoleAssignment struct {
	gorm.Model
	TenantID uint
	UserID   uint
	RoleID   uint
}

type RolePermission struct {
	gorm.Model
	RoleID       uint
	PermissionID uint
}

type Role struct {
	gorm.Model
	Name            string
	RoleAssignments []RoleAssignment
	RolePermissions []RolePermission
}

type Permission struct {
	gorm.Model
	Resource        string
	Action          string
	RolePermissions []RolePermission
}

const (
	RecipeResource = "recipe"
	UserResource   = "user"
	PlanResource   = "plan"
	OrderResource  = "order"
	FoodResource   = "food"
	StapleResource = "staple"
	TenantResource = "tenant"
)

const (
	CreateAction = "create"
	ReadAction   = "read"
	UpdateAction = "update"
	DeleteAction = "delete"
)
