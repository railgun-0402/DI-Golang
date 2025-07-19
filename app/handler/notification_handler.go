package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/railgun-0402/DI-Golang/app/usecase/notification"
)

type NotificationHandler struct {
	Usecase *usecase.NotificationUsecase
}

type EnqueueRequest struct {
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func NewNotificationHandler(u *usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{
		Usecase: u,
	}
}

func (h *NotificationHandler) Enqueue(c echo.Context) error {
	var req EnqueueRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}
	if req.UserID == "" || req.Message == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "user_id and message are required"})
	}
	if err := h.Usecase.Enqueue(c.Request().Context(), req.UserID, req.Title, req.Message); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "queued"})
}
