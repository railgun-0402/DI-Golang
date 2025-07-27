package domain

import (
	"context"
	"time"
)

// メッセージを送る通知情報
type Notification struct {
	ID        string    `dynamodbav:"id"`
	UserID    string    `dynamodbav:"user_id"`
	Title     string    `dynamodbav:"title"`
	Message   string    `dynamodbav:"message"`
	CreatedAt time.Time `dynamodbav:"created_at"`
}

type NotificationRepository interface {
	Save(ctx context.Context, n Notification) error
	// GetAll(ctx context.Context, n Notification) ([]Notification, error)
}

// Publish
type Notifier interface {
	Notify(ctx context.Context, message string) error
}
