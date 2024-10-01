package models

type User struct {
	UserID      int    `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username    string `gorm:"size:50" json:"username" binding:"required"`
	Email       string `gorm:"size:255;unique" json:"email" binding:"required"`
	Password    string `gorm:"size:100" json:"password" binding:"required"`
	PhoneNumber string `gorm:"size:20;unique" json:"phone_number" binding:"required"`
	Picture_URL string `gorm:"size:255" json:"picture_url"`
}