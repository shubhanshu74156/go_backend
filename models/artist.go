package models

import (
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	ArtistID       string `gorm:"type:char(36);index"`
	OrganizationID string
	Name           string
	GrammyCount    int
	Hidden         bool
	Organization   Organization
	Albums         []Album
	Tracks         []Track
}
