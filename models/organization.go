package models

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	OrganizationID string `gorm:"type:varchar(255);index"`
	Name           string
	Users          []User
	Artists        []Artist
	Albums         []Album
	Tracks         []Track
}
