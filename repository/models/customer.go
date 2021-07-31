package repositories

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name        string `gorm:"size:50;unique"`
	DateOfBirth time.Time
	City        string `gorm:"size:20"`
	Zipcode     string `gorm:"size:20"`
	Status      int
}
