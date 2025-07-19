package sqs

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	domain "github.com/railgun-0402/DI-Golang/app/domain/notification"
)

type NotificationQueue struct {
	Client *sqs.Client
	QueueURL string
}

func NewNotificationQueue(c *sqs.Client, url string) *NotificationQueue {
	return &NotificationQueue{
		Client: c,
		QueueURL: url,
	}
}

func (q *NotificationQueue) Enqueue(ctx context.Context, n domain.Notification) error {
	body, _ := json.Marshal(n)
	_, err := q.Client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl: aws.String(q.QueueURL),
		MessageBody: aws.String(string(body)),
	})
	return err
}
