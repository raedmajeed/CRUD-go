package dto

type LoginRequest struct {
	Email			string		`json:"email" gorm:"not null"`
	Password	string		`json:"password" gorm:"not null"`
}


