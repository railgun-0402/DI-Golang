package sqs_repository

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type NotificationQueue struct {
	Client   *sqs.Client
	QueueURL string
}

func NewNotificationQueue(c *sqs.Client, url string) *NotificationQueue {
	return &NotificationQueue{
		Client:   c,
		QueueURL: url,
	}
}
