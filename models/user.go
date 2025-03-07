package models

import (
	"gorm.io/gorm"
)

type UserRole string

const (
	UserRoleAdmin  UserRole = "admin"
	UserRoleEditor UserRole = "editor"
	UserRoleViewer UserRole = "viewer"
)

type User struct {
	gorm.Model
	UserID         string `gorm:"type:char(36);index"`
	OrganizationID string
	Email          string
	Password       string
	Role           UserRole
	Organization   Organization
	Favorites      []Favorite
}
