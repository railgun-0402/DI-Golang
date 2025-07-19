package sqs_repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	domain "github.com/railgun-0402/DI-Golang/app/domain/notification"
)

type Queue struct {
	client   *sqs.Client
	queueURL string
}

func NewQueue(c *sqs.Client, url string) *Queue {
	return &Queue{client: c, queueURL: url}
}

func (q *Queue) ReceiveMessages(ctx context.Context) ([]domain.Message, error) {
	out, err := q.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(q.queueURL),
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     10,
	})
	if err != nil {
		return nil, err
	}

	var msgs []domain.Message
	for _, m := range out.Messages {
		msgs = append(msgs, domain.Message{
			Body:          *m.Body,
			ReceiptHandle: *m.ReceiptHandle,
		})
	}
	return msgs, nil
}

func (q *Queue) DeleteMessage(ctx context.Context, receiptHandle string) error {
	_, err := q.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(q.queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	})
	return err
}

func (q *Queue) Enqueue(ctx context.Context, message string) error {
	_, err := q.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(q.queueURL),
		MessageBody: aws.String(message),
	})
	return err
}
