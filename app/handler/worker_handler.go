package handler

import (
	"context"
	"log"

	usecase "github.com/railgun-0402/DI-Golang/app/usecase/notification"
)

type WorkerHandler struct {
	Usecase *usecase.WorkerUsecase
}

func NewWorkerHandler(uc *usecase.WorkerUsecase) *WorkerHandler {
	return &WorkerHandler{Usecase: uc}
}

func (h *WorkerHandler) Run(ctx context.Context) {
	log.Println("Worker started")
	h.Usecase.Run(ctx)
}
