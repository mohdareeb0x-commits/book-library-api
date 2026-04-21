package models

type Books struct {
	ID            int    `json:"id"`
	Name          string `json:"book"`
	Author        string `json:"author"`
	DatePublished string `json:"date_published"`
	Units         int    `json:"units"`
	Price         int    `json:"price"`
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
