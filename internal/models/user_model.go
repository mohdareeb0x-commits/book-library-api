package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	UserName string `gorm:"not null"`
	Password string `gorm:"not null"`
}
