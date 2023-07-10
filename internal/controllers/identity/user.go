package identity

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/eug/filters"
	"github.com/shinhagunn/todo-backend/internal/helpers"
	"github.com/shinhagunn/todo-backend/internal/models"
	"github.com/shinhagunn/todo-backend/pkg/jwt"
)

var (
	ErrUserPasswordInvalid = helpers.APIError{Code: http.StatusBadRequest, Message: "identity.user.password_invalid"}
	ErrUserNotFound        = helpers.APIError{Code: http.StatusNotFound, Message: "identity.user.not_found"}
	ErrUserJWTGenerate     = helpers.APIError{Code: http.StatusBadRequest, Message: "identity.user.jwt_generate"}
)

// POST: /register
func (h Handler) Register(c *gin.Context) {
	type Payload struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	payload := Payload{}
	if err := h.ParserData(c, &payload, "identity.user"); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user := &models.User{
		Email:    payload.Email,
		Password: payload.Password,
	}

	if err := h.userUsecase.Create(context.TODO(), user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}

// POST: /login
func (h Handler) Login(c *gin.Context) {
	type Payload struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	payload := Payload{}

	if err := h.ParserData(c, &payload, "identity.user"); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userUsecase.First(context.TODO(), filters.WithFieldEqual("email", payload.Email))
	if err != nil || user == nil {
		c.JSON(ErrUserNotFound.Code, ErrUserNotFound.Message)
		return
	}

	if !user.CheckPassword(payload.Password) {
		c.JSON(ErrUserPasswordInvalid.Code, ErrUserPasswordInvalid.Message)
		return
	}

	token, err := jwt.GenerateJWTToken(user.UID)
	if err != nil {
		c.JSON(ErrUserJWTGenerate.Code, ErrUserJWTGenerate.Message)
		return
	}

	c.Header("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
