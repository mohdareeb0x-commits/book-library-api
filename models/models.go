package models

type Books struct {
	ID            int    `form:"id"`
	Name          string `form:"book" binding:"required"`
	Author        string `form:"author" binding:"required"`
	DatePublished string `form:"date_published" binding:"required"`
	Units         int    `form:"units" binding:"required"`
	Price         int    `form:"price" binding:"required"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    Meta        `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	Limit    int `json:"limit,omitempty"`
	Offset      int `json:"offset,omitempty"`
}
