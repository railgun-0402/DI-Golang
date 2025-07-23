package usecase

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	domain "github.com/railgun-0402/DI-Golang/app/domain/notification"
)

type NotificationUsecase struct {
	Queue domain.Queue
	Repo domain.NotificationRepository
}

func NewNotificationUsecase(q domain.Queue, repo domain.NotificationRepository) *NotificationUsecase {
	return &NotificationUsecase{
		Queue: q,
		Repo: repo,
	}
}

// Enqueue はリクエストから受け取った値をキューに送る
func (u *NotificationUsecase) Enqueue(ctx context.Context, userID, title, message string) error {
	n := domain.Notification{
		ID: uuid.New().String(),
		UserID: userID,
		Title: title,
		Message: message,
		CreatedAt: time.Now(),
	}

	body, err := json.Marshal(n)
	if err != nil {
		return err
	}

	err = u.Queue.Enqueue(ctx, string(body))
	if err != nil {
		return err
	}

	return nil
}