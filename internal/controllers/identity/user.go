package identity

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/eug/filters"
	"github.com/shinhagunn/todo-backend/internal/helper"
	"github.com/shinhagunn/todo-backend/internal/models"
	"github.com/shinhagunn/todo-backend/pkg/util"
)

var (
	ErrUserPasswordInvalid = helper.NewAPIError(http.StatusBadRequest, "identity.user.password_invalid")
	ErrUserNotFound        = helper.NewAPIError(http.StatusNotFound, "identity.user.not_found")
	ErrUserJWTGenerate     = helper.NewAPIError(http.StatusInternalServerError, "identity.user.jwt_generate_fail")
)

// POST: /register
func (h Handler) Register(c *gin.Context) {
	type Payload struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	payload := Payload{}
	if ok := h.BindAndValid(c, &payload, "identity.user"); !ok {
		return
	}

	user := &models.User{
		Email:    payload.Email,
		Password: payload.Password,
	}

	if err := h.userUsecase.Create(context.TODO(), user); err != nil {
		h.ResponseError(c, helper.APIError{Code: http.StatusBadRequest, Err: err})
		return
	}

	h.ResponseData(c, http.StatusCreated, user)
}

// POST: /login
func (h Handler) Login(c *gin.Context) {
	type Payload struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	payload := Payload{}

	if ok := h.BindAndValid(c, &payload, "identity.user"); !ok {
		return
	}

	user, err := h.userUsecase.First(context.TODO(), filters.WithFieldEqual("email", payload.Email))
	if err != nil || user == nil {
		h.ResponseError(c, ErrUserNotFound)
		return
	}

	if !user.CheckPassword(payload.Password) {
		h.ResponseError(c, ErrUserPasswordInvalid)
		return
	}

	token, err := util.GenerateToken(user.UID)
	if err != nil {
		h.ResponseError(c, ErrUserJWTGenerate)
		return
	}

	h.ResponseData(c, http.StatusOK, token)
}
