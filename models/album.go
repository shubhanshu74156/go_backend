package models

import (
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	AlbumID        string `gorm:"type:char(36);index"`
	OrganizationID string
	ArtistID       string
	Name           string
	Year           int
	Hidden         bool
	Organization   Organization
	Artist         Artist
	Tracks         []Track
}
