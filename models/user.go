package models

import (
	"github.com/jinzhu/gorm"
	"time"
)
type User struct {
	gorm.Model
	Name string
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Password  string
	CreatedAt       	time.Time
	UpdatedAt       	time.Time
}