package data

import "gorm.io/gorm"

func (UserCredential) TableName() string {
	return "user"
}

type UserCredential struct {
	gorm.Model
	UserName string
	Password string
}

type UserStudent struct {
	gorm.Model
	Active bool
}

type UserAssistant struct {
	gorm.Model
	Active bool
}

type UserInstructor struct {
	gorm.Model
	Active bool
}
