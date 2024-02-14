package database

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string
	Author string
	ISBN   string
}
