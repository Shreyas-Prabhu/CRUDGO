package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name    string
	Company string
}
