package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"size:128;type:varchar(100)"`
	PhoneNum string `gorm:"unique;not null"`
	Password string `gorm:"size:128"`
}
