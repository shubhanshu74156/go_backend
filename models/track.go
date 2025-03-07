package models

import (
	"gorm.io/gorm"
)

type Track struct {
	gorm.Model
	TrackID        string `gorm:"type:char(36);index"`
	OrganizationID string
	AlbumID        string
	ArtistID       string
	Name           string
	Duration       int
	Hidden         bool
	Organization   Organization
	Album          Album
	Artist         Artist
}
