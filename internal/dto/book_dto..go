package dto

import "time"

type CreateBookInput struct {
	Name          string    `json:"name" binding:"required"`
	Author        string    `json:"author" binding:"required"`
	Price         float64   `json:"price" binding:"gte=0"`
	Units         int       `json:"units" binding:"gte=0"`
	DatePublished time.Time `json:"date_published" time_format:"2006-01-02"`
}

type UpdateBookInput struct {
	Name          *string    `json:"name"`
	Author        *string    `json:"author"`
	Price         *float64   `json:"price" binding:"omitempty,gte=0"`
	Units         *int       `json:"units" binding:"omitempty,gte=0"`
	DatePublished *time.Time `json:"date_published" time_format:"2006-01-02"`
}
