package identity

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shinhagunn/eug/filters"
	"github.com/shinhagunn/todo-backend/internal/models"
	"github.com/shinhagunn/todo-backend/pkg/util"
	"gorm.io/gorm"
)

// POST: /register
func (h Handler) Register(c *gin.Context) {
	type Payload struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	payload := Payload{}
	if err := h.BindAndValid(c, &payload); err != nil {
		h.ResponseError(c, http.StatusBadRequest, errors.Wrap(err, "validate params fail"))
		return
	}

	user := &models.User{
		Email:    payload.Email,
		Password: payload.Password,
	}

	if err := h.userUsecase.Create(context.TODO(), user); err != nil {
		h.ResponseError(c, http.StatusInternalServerError, errors.Wrap(err, "register user fail"))
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

	if err := h.BindAndValid(c, &payload); err != nil {
		h.ResponseError(c, http.StatusBadRequest, errors.Wrap(err, "validate params fail"))
		return
	}

	user, err := h.userUsecase.First(context.TODO(), filters.WithFieldEqual("email", payload.Email))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.ResponseError(c, http.StatusNotFound, errors.Wrap(err, "user not found"))
		} else {
			h.ResponseError(c, http.StatusInternalServerError, errors.Wrap(err, "first user fail"))
		}

		return
	}

	if err := user.CheckPassword(payload.Password); err != nil {
		h.ResponseError(c, http.StatusBadRequest, errors.Wrap(err, "check password fail"))
		return
	}

	token, err := util.GenerateToken(user.UID)
	if err != nil {
		h.ResponseError(c, http.StatusInternalServerError, errors.Wrap(err, "generate token fail"))
		return
	}

	h.ResponseData(c, http.StatusOK, token)
}
