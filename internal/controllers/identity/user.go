package identity

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/eug/filters"
	"github.com/shinhagunn/todo-backend/internal/models"
	"github.com/shinhagunn/todo-backend/pkg/jwt"
)

// TODO: Add support handle print error

type ErrorCustom struct {
	Code int
	Mess string
}

var (
	ErrUserPasswordInvalid = ErrorCustom{http.StatusBadRequest, "identity.user.password_invalid"}
	ErrUserNotFound        = ErrorCustom{http.StatusNotFound, "identity.user.not_found"}
	ErrUserJWTGenerate     = ErrorCustom{http.StatusBadRequest, "identity.user.jwt_generate"}
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
		c.JSON(ErrUserNotFound.Code, ErrUserNotFound.Mess)
		return
	}

	if !user.CheckPassword(payload.Password) {
		c.JSON(ErrUserPasswordInvalid.Code, ErrUserPasswordInvalid.Mess)
		return
	}

	token, err := jwt.GenerateJWTToken(user.UID)
	if err != nil {
		c.JSON(ErrUserJWTGenerate.Code, ErrUserJWTGenerate.Mess)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
