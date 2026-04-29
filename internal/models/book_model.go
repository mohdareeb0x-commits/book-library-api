package models

import "time"

type Book struct {
    ID            uint      `gorm:"primaryKey"`
    Name          string    `gorm:"not null"`
    Author        string    `gorm:"not null"`
    Price         float64
    Units         int
    DatePublished time.Time `time_format:"2006-01-02"`
    CreatedAt     time.Time
    UpdatedAt     time.Time
}

