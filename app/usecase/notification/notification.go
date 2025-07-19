package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	domain "github.com/railgun-0402/DI-Golang/app/domain/notification"
	sqs "github.com/railgun-0402/DI-Golang/app/infra/aws/sqs/notification"
)

type NotificationUsecase struct {
	sqsQueue *sqs.NotificationQueue
}

func NewNotificationUsecase(q *sqs.NotificationQueue) *NotificationUsecase {
	return &NotificationUsecase{sqsQueue: q}
}

func (u *NotificationUsecase) Enqueue(ctx context.Context, userID, title, message string) error {
	n := domain.Notification{
		ID: uuid.New().String(),
		UserID: userID,
		Title: title,
		Message: message,
		CreatedAt: time.Now(),
	}
	return u.sqsQueue.Enqueue(ctx, n)
}