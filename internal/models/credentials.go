package models

import "gorm.io/gorm"

type Credential struct {
	gorm.Model
	ApiKey  string `json:"apiKey"`
	AppName string `gorm:"unique" json:"appName"`
}
