package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Name string `json:"Name"`
	Link string `json:"NewLink"`
	RedirectTO string `json:"RedirectTO"`
}

