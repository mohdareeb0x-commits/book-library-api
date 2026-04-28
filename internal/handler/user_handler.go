package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/internal/dto"
	"github.com/mohdareeb0x-commits/book-library-api/internal/response"
	"github.com/mohdareeb0x-commits/book-library-api/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) CreateUser(c *gin.Context) {
	var register dto.RegisterInput

	if err := c.ShouldBindJSON(&register); err != nil {
		response.Fail(c, http.StatusBadRequest, "INVALID_USER_CREDENTIAL", "unable to register beacause of invalid credentials")
		return
	}

	user, err := h.authService.CreateUser(register)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, "USER_REGISTERATION_FAILED", "failed user registeration in database")
		return
	}

	response.OK(c, user, nil)
}