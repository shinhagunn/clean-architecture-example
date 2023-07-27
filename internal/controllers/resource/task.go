package resource

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shinhagunn/eug/filters"
	"github.com/shinhagunn/todo-backend/internal/models"
	"gorm.io/gorm"
)

// TODO: Add support func GET include: page, limit, total, offset, order

// GET: /tasks
func (h Handler) GetTasks(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	type Payload struct {
		Page  int   `form:"page" json:"page"`
		Limit int   `form:"limit" json:"limit"`
		Total int64 `form:"-" json:"total"`
	}

	payload := Payload{}
	if err := h.BindAndValid(c, &payload); err != nil {
		h.ResponseError(c, http.StatusBadRequest, errors.Wrap(err, "validate params fail"))
		return
	}

	if payload.Page <= 0 {
		payload.Page = 1
	}

	if payload.Limit <= 0 {
		payload.Limit = 10
	}

	offset := (payload.Page - 1) * payload.Limit

	ctx := context.TODO()
	tasks := h.taskUsecase.Find(
		ctx,
		// TODO: Add support filters WithCount
		filters.WithOffset(offset),
		filters.WithFieldEqual("user_id", user.ID),
		filters.WithFieldEqual("status", models.TaskStatusProcessing),
	)

	h.ResponseData(c, http.StatusOK, tasks)
}

// TODO: Add support deadline_at
// POST: /tasks
func (h Handler) CreateTask(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	type Payload struct {
		CategoryID int64  `form:"category_id" json:"category_id" binding:"required"`
		Level      int64  `form:"level" json:"level" binding:"required"`
		Name       string `from:"name" json:"name" binding:"required"`
	}

	payload := Payload{}
	if err := h.BindAndValid(c, &payload); err != nil {
		h.ResponseError(c, http.StatusBadRequest, errors.Wrap(err, "validate params fail"))
		return
	}

	task := &models.Task{
		UserID:     user.ID,
		CategoryID: payload.CategoryID,
		Level:      payload.Level,
		Name:       payload.Name,
	}

	ctx := context.TODO()
	if err := h.taskUsecase.Create(ctx, task); err != nil {
		h.ResponseError(c, http.StatusInternalServerError, errors.Wrap(err, "create task fail"))
		return
	}

	h.ResponseData(c, http.StatusCreated, task)
}

// PUT: /tasks
func (h Handler) UpdateTask(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	type Payload struct {
		ID         int64  `form:"id" json:"id" binding:"required"`
		CategoryID int64  `form:"category_id" json:"category_id"`
		Level      int64  `form:"level" json:"level"`
		Name       string `from:"name" json:"name"`
	}

	payload := Payload{}
	if err := h.BindAndValid(c, &payload); err != nil {
		h.ResponseError(c, http.StatusBadRequest, errors.Wrap(err, "validate params fail"))
		return
	}

	ctx := context.TODO()
	task, err := h.taskUsecase.First(
		ctx,
		filters.WithFieldEqual("user_id", user.ID),
		filters.WithFieldEqual("id", payload.ID),
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.ResponseError(c, http.StatusNotFound, errors.Wrap(err, "task not found"))
		} else {
			h.ResponseError(c, http.StatusInternalServerError, errors.Wrap(err, "first task fail"))
		}

		return
	}

	taskUpdates := &models.Task{}
	if payload.CategoryID > 0 {
		taskUpdates.CategoryID = payload.CategoryID
	}

	if payload.Level > 0 {
		taskUpdates.Level = payload.Level
	}

	if len(payload.Name) > 0 {
		taskUpdates.Name = payload.Name
	}

	if err := h.taskUsecase.Updates(ctx, task, taskUpdates); err != nil {
		h.ResponseError(c, http.StatusInternalServerError, errors.Wrap(err, "updates user fail"))
		return
	}

	h.ResponseData(c, http.StatusOK, task)
}

// DELETE /task
func (h Handler) DeleteTask(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	type Payload struct {
		ID int64 `form:"id" json:"id" binding:"required"`
	}

	payload := Payload{}
	if err := h.BindAndValid(c, &payload); err != nil {
		h.ResponseError(c, http.StatusBadRequest, errors.Wrap(err, "validate params fail"))
		return
	}

	ctx := context.TODO()
	task, err := h.taskUsecase.First(
		ctx,
		filters.WithFieldEqual("user_id", user.ID),
		filters.WithFieldEqual("id", payload.ID),
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.ResponseError(c, http.StatusNotFound, errors.Wrap(err, "task not found"))
		} else {
			h.ResponseError(c, http.StatusInternalServerError, errors.Wrap(err, "first task fail"))
		}

		return
	}

	if err := h.taskUsecase.Delete(ctx, task); err != nil {
		h.ResponseError(c, http.StatusInternalServerError, errors.Wrap(err, "delete task fail"))
		return
	}

	c.JSON(http.StatusOK, 200)
}
