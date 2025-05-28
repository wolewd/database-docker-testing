package models

type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}
