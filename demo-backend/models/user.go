package models

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required" gorm:"unique"`
}
