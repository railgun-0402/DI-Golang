package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/railgun-0402/DI-Golang/app/usecase/device"
)

type DeviceHandler struct {
	Usecase *usecase.DeviceUsecase
}

type RegisterDeviceRequest struct {
	UserID   string `json:"user_id"`
    Token    string `json:"token"`
    Platform string `json:"platform"` // ios or android
}

func (h *DeviceHandler) Register(c echo.Context) error {
	var req RegisterDeviceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.Usecase.Register(req.UserID, req.Token, req.Platform); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "ok"})
}