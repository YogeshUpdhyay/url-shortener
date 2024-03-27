package models

import "gorm.io/gorm"

type ShortUrl struct {
	gorm.Model
	Slug string `gorm:"primaryKey" json:"slug"`
	URL  string `json:"url"`
}
