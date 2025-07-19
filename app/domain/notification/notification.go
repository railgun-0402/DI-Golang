package domain

import (
	"context"
	"time"
)

type Notification struct {
	ID        string    `dynamodbav:"id"`
	UserID    string    `dynamodbav:"user_id"`
	Title     string    `dynamodbav:"title"`
	Message   string    `dynamodbav:"message"`
	CreatedAt time.Time `dynamodbav:"created_at"`
}

type NotificationRepository interface {
	Save(ctx context.Context, n Notification) error
}
