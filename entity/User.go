package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email 						string			`json:"email" gorm:"unique; not null"`
	Password 					string			`json:"password" gorm:"not null"`
	Role							string			`json:"role" gorm:"default: 'user'"`
	IsAccountLocked		bool 				`json:"is_account_locked" gorm:"default:false"`
	PhoneNumber				string			`json:"phone_number"`
}