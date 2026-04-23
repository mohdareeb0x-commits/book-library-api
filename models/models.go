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
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
