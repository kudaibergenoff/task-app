package domain

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title  string `gorm:"size:255;index:idx_title, unique"`
	Status bool   `gorm:"default:true"`
}
