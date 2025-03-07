package models

import (
	"gorm.io/gorm"
)

type FavoriteCategory string

const (
	FavoriteCategoryArtist FavoriteCategory = "artist"
	FavoriteCategoryAlbum  FavoriteCategory = "album"
	FavoriteCategoryTrack  FavoriteCategory = "track"
)

type Favorite struct {
	gorm.Model
	FavoriteID string `gorm:"type:char(36);index"`
	UserID     string
	Category   FavoriteCategory
	ItemID     string
	Hidden     bool
	User       User
}
